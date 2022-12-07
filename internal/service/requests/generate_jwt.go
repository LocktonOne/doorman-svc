package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	"gitlab.com/tokene/doorman/resources"
)

type Claims struct {
	resources.JwtClaims
}

func NewGenerateJwt(r *http.Request) (resources.JwtClaimsAttributes, error) {
	var request Claims
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request.Attributes, errors.Wrap(err, "failed to unmarshal")
	}
	return request.Attributes, validateClaims(request)
}
func validateClaims(r Claims) error {
	return validation.Errors{
		"/attributes/eth_address": validation.Validate(&r.Attributes.EthAddress, validation.Required, validation.Match(helpers.AddressRegexp)),
		"/attributes/purpose": validation.Validate(string(r.Attributes.Purpose.Type),
			validation.Required,
			validation.By(helpers.ValidatePurposes)),
		"/type": validation.Validate(&r.Type,
			validation.Required),
	}.Filter()
}
