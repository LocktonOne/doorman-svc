package connector

import (
	"bytes"
	"encoding/json"
	"io"
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
func (c Connector) DoAuthRequest(method string, url string, token string, body []byte) (*http.Response, error) {
	postBody := bytes.NewBuffer(body)

	req, err := http.NewRequest(method, url, postBody)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	req.Header.Set("Authorization", "Bearer "+token)

	response, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}

	return response, nil
}
func (c Connector) ParseModel(body io.Reader, model interface{}) error {
	if err := json.NewDecoder(body).Decode(model); err != nil {
		return errors.Wrap(err, "failed to unmarshal")
	}
	return nil
}
func (c Connector) GenerateJwtPair(address string, purpose string) (resources.JwtPairResponse, error) {

	postBody, err := json.Marshal(NewClaimsModel(address, purpose))
	if err != nil {
		return resources.JwtPairResponse{}, errors.Wrap(err, "failed to marshal")
	}

	response, err := c.DoAuthRequest("POST", c.ServiceUrl+"/get_token_pair", "", postBody)
	if err != nil {
		return resources.JwtPairResponse{}, err
	}
	defer response.Body.Close()

	model := &resources.JwtPairResponse{}
	return *model, c.ParseModel(response.Body, model)
}

func (c Connector) ValidateJwt(token string) (string, error) {
	response, err := c.DoAuthRequest("POST", c.ServiceUrl+"/validate_token", token, nil)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	model := &resources.JwtValidationResponse{}
	return *&model.Data.Attributes.EthAddress, c.ParseModel(response.Body, model)
}

func (c Connector) RefreshJwt(refreshToken string) (resources.JwtPairResponse, error) {
	response, err := c.DoAuthRequest("POST", c.ServiceUrl+"/refresh_token", refreshToken, nil)
	if err != nil {
		return resources.JwtPairResponse{}, err
	}
	defer response.Body.Close()

	model := &resources.JwtPairResponse{}
	return *model, c.ParseModel(response.Body, model)
}

func (c Connector) GetAuthToken(r *http.Request) (string, error) {
	return helpers.Authenticate(r)
}

func (c Connector) CheckPermission(owner string, token string) (bool, error) {
	postBody, err := json.Marshal(NewCheckPermissionModel(owner))
	if err != nil {
		return false, errors.Wrap(err, "failed to marshal")
	}

	response, err := c.DoAuthRequest("POST", c.ServiceUrl+"/check_permission", token, postBody)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	return response.StatusCode == http.StatusNoContent, nil
}
