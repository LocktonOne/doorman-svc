package requests

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strings"
)

type GetPermissionRequest struct {
	Id       string
	Resource string
}

func NewGetPermissionRequest(r *http.Request) (GetPermissionRequest, error) {
	request := GetPermissionRequest{}
	request.Id = strings.ToUpper(r.URL.Query().Get("id"))
	request.Resource = strings.ToUpper(r.URL.Query().Get("resource"))
	if request.Resource == "" {
		request.Resource = "*"
	}
	err := request.validate(request.Id)
	if err != nil {
		return GetPermissionRequest{}, err
	}
	return request, nil
}

func (p GetPermissionRequest) validate(id string) error {
	var permissions = []string{ //todo make better
		"CREATE", "VIEW", "READ",
	}

	for _, permission := range permissions {
		if permission == id {
			return nil
		}
	}
	return errors.New("invalid permission id")
}
