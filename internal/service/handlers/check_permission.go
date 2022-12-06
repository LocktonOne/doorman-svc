package handlers

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	"gitlab.com/tokene/doorman/internal/service/requests"
	"net/http"
)

func CheckPermission(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	permission, err := requests.NewGetPermissionRequest(r)
	if err != nil {
		logger.WithError(err).Debug(err)
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	_, address, err := helpers.GetAccessTokenInfo(r)
	if err != nil {
		logger.WithError(err).Debug("failed to retrieve access token")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	if permission.Id != "" {
		accManager, err := helpers.GetAddressAccManagement(r)
		if err != nil {
			logger.WithError(err).Debug("can't get contract address")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		success, err := helpers.CheckPermissionsByAddress(accManager, common.HexToAddress(address), helpers.Cfg.GetClient(), permission.Id, helpers.AllResource)
		if err != nil {
			logger.WithError(err).Debug("failed to check account permissions")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if !success {
			logger.WithError(err).Debug("user has no rights to get resource")
			ape.RenderErr(w, problems.Forbidden())
			return

		}

		w.WriteHeader(http.StatusOK)

	}

	w.WriteHeader(http.StatusNoContent)
}
