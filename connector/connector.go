package connector

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

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

func (c Connector) GenerateJwtPair(claims resources.JwtClaims) (resources.JwtPairResponse, error) {
	postBody, _ := json.Marshal(claims)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(c.ServiceUrl+"/get_token_pair", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	var request resources.JwtPairResponse
	if err := json.NewDecoder(resp.Body).Decode(&request); err != nil {
		return resources.JwtPairResponse{}, errors.Wrap(err, "failed to unmarshal")
	}

	return request, nil
}
func (c Connector) ValidateJwt(token string) (bool, error) {
	return false, nil
}
func (c Connector) RefreshJwt(refreshToken string) (resources.JwtPairResponse, error) {
	return resources.JwtPair{}, nil
}
