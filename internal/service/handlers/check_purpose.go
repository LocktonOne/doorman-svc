package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	"gitlab.com/tokene/doorman/resources"
)

func CheckPurpose(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	token, err := helpers.Authenticate(r)
	if err != nil {
		logger.WithError(err).Debug("failed to get token")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	purpose, _, err := helpers.RetrieveJwtToken(token, r)
	if err != nil {
		logger.WithError(err).Debug("failed to retrieve token")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	err = helpers.ValidatePurposes(purpose)
	if err != nil {
		logger.WithError(err).Debug("unknown purpose")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	result := resources.Purpose{
		Type: purpose,
	}

	ape.Render(w, result)
}
