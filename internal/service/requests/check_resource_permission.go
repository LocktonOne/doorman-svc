package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	"gitlab.com/tokene/doorman/resources"
)

func NewCheckResourcePermission(r *http.Request) (resources.CheckResourcePermissionAttributes, error) {
	var request resources.CheckResourcePermission
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request.Attributes, errors.Wrap(err, "failed to unmarshal")
	}
	return request.Attributes, validateCheckResourcePermission(request.Attributes)
}
func validateCheckResourcePermission(r resources.CheckResourcePermissionAttributes) error {
	return validation.Errors{
		"/attributes/owner": validation.Validate(&r.Owner, validation.Required, validation.Match(helpers.AddressRegexp)),
	}.Filter()
}
