package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	"gitlab.com/tokene/doorman/resources"
)

func NewValidateToken(r *http.Request) (resources.JwtValidationAttributes, error) {
	var request resources.JwtValidation
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request.Attributes, errors.Wrap(err, "failed to unmarshal")
	}
	return request.Attributes, validateJwtValidationAttributes(request)
}
func validateJwtValidationAttributes(r resources.JwtValidation) error {
	return validation.Errors{
		"/attributes/eth_address": validation.Validate(&r.Attributes.EthAddress, validation.Required, validation.Match(helpers.AddressRegexp)),
		"/type":                   validation.Validate(&r.Type, validation.Required),
	}.Filter()
}
