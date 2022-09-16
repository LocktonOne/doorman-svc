package connector

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/tokene/doorman/resources"
)

type ConnectorMockKyc struct {
	ServiceUrl string
}

func NewConnectorMockKyc(serviceUrl string) ConnectorI {
	return ConnectorMockKyc{
		ServiceUrl: serviceUrl,
	}
}

func (c ConnectorMockKyc) GenerateJwtPair(address string, purpose string) (resources.JwtPairResponse, error) {
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

func (c ConnectorMockKyc) ValidateJwt(token string) (string, error) {
	return strings.ToLower("0x750Bd531CEA1f68418DDF2373193CfbD86A69058"), nil
}

func (c ConnectorMockKyc) RefreshJwt(refreshToken string) (resources.JwtPairResponse, error) {
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
func (c ConnectorMockKyc) GetAuthToken(r *http.Request) (string, error) {
	return "token", nil
}
