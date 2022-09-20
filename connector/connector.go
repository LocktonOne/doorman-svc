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
func (c Connector) DoAuthRequest(method string, url string, token string, body interface{}, parseResponseModel interface{}) (*http.Response, error) {
	postBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	req.Header.Set("Authorization", "Bearer "+token)

	response, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}

	return response, c.ParseModel(response.Body, parseResponseModel)
}
func (c Connector) ParseModel(body io.Reader, model interface{}) error {
	if model == nil {
		return nil
	}
	if err := json.NewDecoder(body).Decode(model); err != nil {
		return errors.Wrap(err, "failed to unmarshal")
	}
	return nil
}
func (c Connector) GenerateJwtPair(address string, purpose string) (resources.JwtPairResponse, error) {
	model := &resources.JwtPairResponse{}

	response, err := c.DoAuthRequest("POST", c.ServiceUrl+"/get_token_pair", "", NewClaimsModel(address, purpose), model)
	defer response.Body.Close()

	return *model, err
}

func (c Connector) ValidateJwt(token string) (string, error) {
	model := &resources.JwtValidationResponse{}

	response, err := c.DoAuthRequest("POST", c.ServiceUrl+"/validate_token", token, nil, model)
	defer response.Body.Close()

	return *&model.Data.Attributes.EthAddress, err
}

func (c Connector) RefreshJwt(refreshToken string) (resources.JwtPairResponse, error) {
	model := &resources.JwtPairResponse{}

	response, err := c.DoAuthRequest("POST", c.ServiceUrl+"/refresh_token", refreshToken, nil, model)
	defer response.Body.Close()

	return *model, err
}

func (c Connector) GetAuthToken(r *http.Request) (string, error) {
	return helpers.Authenticate(r)
}

func (c Connector) CheckPermission(owner string, token string) (bool, error) {
	response, err := c.DoAuthRequest("POST", c.ServiceUrl+"/check_permission", token, NewCheckPermissionModel(owner), nil)
	defer response.Body.Close()

	return response.StatusCode == http.StatusNoContent, err
}
