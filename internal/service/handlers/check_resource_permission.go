package handlers

import (
	"github.com/ethereum/go-ethereum/common"
	"net/http"
	"strings"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	"gitlab.com/tokene/doorman/internal/service/requests"
)

func CheckResourcePermission(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	owner, err := requests.NewCheckResourcePermission(r)
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
	client, err := helpers.EthRPCConfig(r).EthClient()
	if err != nil {
		logger.WithError(err).Debug("failed connection to eth ")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if address != strings.ToLower(owner) {
		success, err := helpers.CheckPermissionsByAddress(helpers.RegistryConfig(r).Address, common.HexToAddress(address), client)
		if err != nil {
			logger.WithError(err).Debug("Internal error")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if !success {
			logger.WithError(err).Debug("user has no rights to get resource")
			ape.RenderErr(w, problems.Forbidden())
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
