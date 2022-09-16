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
		Key:        resources.Key{Type: resources.VALIDATE_TOKEN},
		Attributes: resources.JwtValidationAttributes{EthAddress: address},
	}
	return model
}
func NewCheckPermissionModel(owner string) resources.CheckResourcePermission {
	model := resources.CheckResourcePermission{
		Key:        resources.Key{Type: resources.CHECK_RESOURCE_PERMISSION},
		Attributes: resources.CheckResourcePermissionAttributes{Owner: owner},
	}
	return model
}
