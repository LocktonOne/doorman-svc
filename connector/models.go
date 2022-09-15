package connector

import (
	"gitlab.com/tokene/doorman/resources"
)

func NewClaimsModel(address string, purpose string) resources.JwtClaims {
	model := resources.JwtClaims{
		Key: resources.Key{Type: resources.JWT_CLAIMS},
		Attributes: resources.JwtClaimsAttributes{
			EthAddress: address,
			Purpose: resources.Purpose{
				Type: purpose,
			},
		},
	}
	return model
}
func NewJwtValidationModel(address string) resources.JwtValidation {
	model := resources.JwtValidation{
		Key:        resources.Key{Type: "validate_token"},
		Attributes: resources.JwtValidationAttributes{EthAddress: address},
	}
	return model
}
