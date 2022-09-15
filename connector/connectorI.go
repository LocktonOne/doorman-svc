package connector

import (
	"gitlab.com/tokene/doorman/resources"
)

type ConnectorI interface {
	GenerateJwtPair(model resources.JwtClaims) (resources.JwtPairResponse, error)
	ValidateJwt(token string) (bool, error)
	RefreshJwt(refreshToken string) (resources.JwtPairResponse, error)
}
