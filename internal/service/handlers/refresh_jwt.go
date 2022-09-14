package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	"gitlab.com/tokene/doorman/resources"
)

func RefreshJwt(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	token, err := helpers.Authenticate(r)
	if err != nil {
		logger.WithError(err).Debug(err)
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	address, err := helpers.RetrieveToken(token, r)
	if err != nil {
		logger.WithError(err).Debug("failed to retrieve refresh token")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// success logic
	sessionToken, sessioExp, err := helpers.GenerateJWT(
		resources.JwtClaimsAttributes{
			EthAddress: address,
			Purpose:    resources.Purpose{Type: string(resources.SESSION_JWT)},
		}, helpers.ServiceConfig(r))
	if err != nil {
		logger.WithError(err).Debug(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	refreshToken, refreshExp, err := helpers.GenerateRefreshToken(
		resources.JwtClaimsAttributes{
			EthAddress: address,
			Purpose:    resources.Purpose{Type: string(resources.SESSION_JWT)},
		}, helpers.ServiceConfig(r))
	if err != nil {
		logger.WithError(err).Debug(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := NewJwtPairResponseModel(
		NewJwtModel(sessionToken, string(resources.SESSION_JWT), sessioExp),
		NewJwtModel(refreshToken, string(resources.REFRESH_JWT), refreshExp),
	)
	ape.Render(w, response)
}
