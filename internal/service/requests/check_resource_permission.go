package requests

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"gitlab.com/tokene/doorman/internal/service/helpers"
)

func NewCheckResourcePermission(r *http.Request) (string, error) {
	owner := r.URL.Query().Get("owner")
	if err := validation.Validate(owner, validation.Required, validation.Match(helpers.AddressRegexp)); err != nil {
		return owner, errors.Wrap(err, "failed to get parameter")
	}
	return owner, nil
}
