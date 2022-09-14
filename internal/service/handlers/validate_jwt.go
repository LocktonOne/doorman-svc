package handlers

import (
	"net/http"
	"strings"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	"gitlab.com/tokene/doorman/internal/service/requests"
)

func ValidateJWT(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	req, err := requests.NewValidateToken(r)
	if err != nil {
		logger.WithError(err).Debug(err)
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	token, err := helpers.Authenticate(r)
	if err != nil {
		logger.WithError(err).Debug(err)
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	address, err := helpers.RetrieveToken(token, r)
	if err != nil {
		logger.WithError(err).Debug("failed to retrieve token")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	if strings.ToLower(address) != strings.ToLower(req.EthAddress) {
		logger.WithError(err).Debug("address is not the same")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}
	ape.Render(w, http.StatusOK)
}
