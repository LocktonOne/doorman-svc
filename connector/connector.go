package connector

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	"gitlab.com/tokene/doorman/resources"
)

type Connector struct {
	ServiceUrl string
	Client     *http.Client
}

func NewConnector(serviceUrl string) ConnectorI {
	return Connector{
		ServiceUrl: serviceUrl,
		Client: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

func (c Connector) GenerateJwtPair(address string, purpose string) (resources.JwtPairResponse, error) {

	postBody, err := json.Marshal(NewClaimsModel(address, purpose))
	if err != nil {
		return resources.JwtPairResponse{}, errors.Wrap(err, "failed to marshal")
	}

	req, err := http.NewRequest("POST", c.ServiceUrl+"/get_token_pair", bytes.NewBuffer(postBody))
	if err != nil {
		return resources.JwtPairResponse{}, errors.Wrap(err, "failed to make request")
	}

	response, err := c.Client.Do(req)
	if err != nil {
		return resources.JwtPairResponse{}, errors.Wrap(err, "failed to request")
	}

	defer response.Body.Close()

	var request resources.JwtPairResponse
	if err := json.NewDecoder(response.Body).Decode(&request); err != nil {
		return resources.JwtPairResponse{}, errors.Wrap(err, "failed to unmarshal")
	}

	return request, nil
}

func (c Connector) ValidateJwt(token string) (string, error) {

	req, err := http.NewRequest("POST", c.ServiceUrl+"/validate_token", nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to make request")
	}

	req.Header.Set("Authorization", "Bearer "+token)

	response, err := c.Client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "failed to make request")
	}
	defer response.Body.Close()

	var request resources.JwtValidationResponse
	if err := json.NewDecoder(response.Body).Decode(&request); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal")
	}

	return request.Data.Attributes.EthAddress, nil
}

func (c Connector) RefreshJwt(refreshToken string) (resources.JwtPairResponse, error) {

	req, err := http.NewRequest("POST", c.ServiceUrl+"/refresh_token", nil)
	if err != nil {
		return resources.JwtPairResponse{}, errors.Wrap(err, "failed to create request")
	}

	req.Header.Set("Authorization", "Bearer "+refreshToken)

	response, err := c.Client.Do(req)
	if err != nil {
		return resources.JwtPairResponse{}, errors.Wrap(err, "failed to make request")
	}

	defer response.Body.Close()

	var request resources.JwtPairResponse
	if err := json.NewDecoder(response.Body).Decode(&request); err != nil {
		return resources.JwtPairResponse{}, errors.Wrap(err, "failed to unmarshal")
	}

	return request, nil
}

func (c Connector) GetAuthToken(r *http.Request) (string, error) {
	return helpers.Authenticate(r)
}

func (c Connector) CheckPermission(owner string, token string) (bool, error) {

	postBody, err := json.Marshal(NewCheckPermissionModel(owner))
	if err != nil {
		return false, errors.Wrap(err, "failed to marshal")
	}

	req, err := http.NewRequest("POST", c.ServiceUrl+"/check_permission", bytes.NewBuffer(postBody))
	if err != nil {
		return false, errors.Wrap(err, "failed to create request")
	}

	req.Header.Set("Authorization", "Bearer "+token)

	response, err := c.Client.Do(req)
	if err != nil {
		return false, errors.Wrap(err, "failed to make request")
	}

	defer response.Body.Close()

	return response.Status == "204 No Content", nil
}
