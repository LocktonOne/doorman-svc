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
}

func NewConnector(serviceUrl string) ConnectorI {
	return Connector{
		ServiceUrl: serviceUrl,
	}
}

func (c Connector) GenerateJwtPair(address string, purpose string) (resources.JwtPairResponse, error) {
	postBody, err := json.Marshal(NewClaimsModel(address, purpose))
	if err != nil {
		return resources.JwtPairResponse{}, errors.Wrap(err, "failed to marshal")
	}

	responseBody := bytes.NewBuffer(postBody)

	response, err := http.Post(c.ServiceUrl+"/get_token_pair", "application/json", responseBody)

	if err != nil || response.StatusCode != http.StatusOK {
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
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", c.ServiceUrl+"/validate_token", nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	var request resources.JwtValidationResponse
	if err := json.NewDecoder(response.Body).Decode(&request); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal")
	}
	return request.Data.Attributes.EthAddress, nil
}

func (c Connector) RefreshJwt(refreshToken string) (resources.JwtPairResponse, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", c.ServiceUrl+"/refresh_token", nil)

	if err != nil {
		return resources.JwtPairResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+refreshToken)

	response, err := client.Do(req)
	if err != nil {
		return resources.JwtPairResponse{}, err
	}
	if response.Status != "200 OK" {
		return resources.JwtPairResponse{}, errors.New("bad status")
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
	if req.Response.Status != "200 OK" {
		return false, nil
	}

	return true, nil
}
