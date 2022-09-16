package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	"gitlab.com/tokene/doorman/resources"
)

func ValidateJWT(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

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
	result := resources.JwtValidationResponse{
		Data: resources.JwtValidation{
			Key: resources.Key{Type: resources.VALIDATE_TOKEN},
			Attributes: resources.JwtValidationAttributes{
				EthAddress: address,
			},
		},
	}
	ape.Render(w, result)
}
