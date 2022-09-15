package connector

import (
	"net/http"

	"gitlab.com/tokene/doorman/resources"
)

type ConnectorI interface {
	GenerateJwtPair(address string, purpose string) (resources.JwtPairResponse, error)
	ValidateJwt(token string, address string) (bool, error)
	RefreshJwt(refreshToken string) (resources.JwtPairResponse, error)
	GetAuthToken(r *http.Request) (string, error)
}
