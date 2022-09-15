package connector

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
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
	postBody, _ := json.Marshal(NewClaimsModel(address, purpose))
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(c.ServiceUrl+"/get_token_pair", "application/json", responseBody)

	if err != nil || resp.StatusCode != http.StatusOK {
		return resources.JwtPairResponse{}, errors.Wrap(err, "failed to request")
	}

	defer resp.Body.Close()

	var request resources.JwtPairResponse
	if err := json.NewDecoder(resp.Body).Decode(&request); err != nil {
		return resources.JwtPairResponse{}, errors.Wrap(err, "failed to unmarshal")
	}

	return request, nil
}

func (c Connector) ValidateJwt(token string, address string) (bool, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	postBody, _ := json.Marshal(NewJwtValidationModel(address))

	req, err := http.NewRequest("POST", c.ServiceUrl+"/validate_token", nil)
	req.Body.Read(postBody)

	if err != nil {
		return false, err
	}

	req.Header.Set("Bearer", token)

	response, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	return response.StatusCode == http.StatusOK, nil
}

func (c Connector) RefreshJwt(refreshToken string) (resources.JwtPairResponse, error) {
	return resources.JwtPairResponse{}, nil
}
