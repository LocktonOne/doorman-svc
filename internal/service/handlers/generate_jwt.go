package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	"gitlab.com/tokene/doorman/internal/service/requests"
	"gitlab.com/tokene/doorman/resources"
)

func NewJwtModel(token string, tokenType string, expTime int64) resources.Jwt {
	model := resources.Jwt{
		Key:        resources.Key{ID: token, Type: resources.ResourceType(tokenType)},
		Attributes: resources.JwtAttributes{ExpiresIn: expTime},
	}
	return model
}
func NewJwtPairResponseModel(sessionToken resources.Jwt, refreshToken resources.Jwt) resources.JwtPairResponse {
	model := resources.JwtPairResponse{
		Data: resources.JwtPair{
			Key: resources.Key{Type: resources.JWT_PAIR},
			Attributes: resources.JwtPairAttributes{
				SessionToken: sessionToken,
				RefreshToken: refreshToken,
			},
		},
	}
	return model
}
func GenerateJwtPair(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewGenerateJwt(r)
	if err != nil {
		logger.WithError(err).Debug("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	sessionToken, sessioExp, err := helpers.GenerateJWT(request, helpers.ServiceConfig(r))
	if err != nil {
		logger.WithError(err).Debug(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	refreshToken, refreshExp, err := helpers.GenerateRefreshToken(request, helpers.ServiceConfig(r))
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
