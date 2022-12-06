package requests

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strings"
)

type GetPermissionRequest struct {
	Id string
}

func NewGetPermissionRequest(r *http.Request) (GetPermissionRequest, error) {
	request := GetPermissionRequest{}
	request.Id = strings.ToUpper(r.URL.Query().Get("id"))
	err := validate(request.Id)
	if err != nil {
		return GetPermissionRequest{}, err
	}
	return request, nil
}

var permissions = []string{ //todo make better
	"CREATE", "VIEW", "READ",
}

func validate(id string) error {
	for _, permission := range permissions {
		if permission == id {
			return nil
		}
	}
	return errors.New("invalid permission id")
}
