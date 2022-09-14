package helpers

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"gitlab.com/tokene/doorman/internal/config"
	"gitlab.com/tokene/doorman/resources"
)

type standardClaims struct {
	Address string `json:"address"`
	jwt.StandardClaims
}
type refreshTokenClaims struct {
	Address string `json:"address"`
	jwt.StandardClaims
}

func GenerateJWT(user_claims resources.JwtClaimsAttributes, cfg *config.ServiceConfig) (string, int64, error) {
	expirationTime := time.Now().Add(cfg.TokenExpireTime)
	claims := &standardClaims{
		Address: strings.ToLower(user_claims.EthAddress),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.TokenKey))
	if err != nil {
		return "", expirationTime.Unix(), errors.Wrap(err, "failed to generate JWT")
	}
	return tokenString, expirationTime.Unix(), nil
}

func GenerateRefreshToken(user_claims resources.JwtClaimsAttributes, cfg *config.ServiceConfig) (string, int64, error) {
	expirationTime := time.Now().Add(cfg.RefreshTokenExpireTime)
	claims := refreshTokenClaims{
		Address: strings.ToLower(user_claims.EthAddress),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.TokenKey))
	if err != nil {
		return "", expirationTime.Unix(), errors.Wrap(err, "failed to generate refresh token")
	}
	return refreshTokenString, expirationTime.Unix(), nil
}

func parseStandardJWT(tokenString string, r *http.Request) (*standardClaims, error) {
	claims := &standardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ServiceConfig(r).TokenKey), nil
	})
	if err != nil {
		return claims, err
	}
	if !token.Valid {
		return claims, errors.New("invalid token")
	}
	return claims, nil
}

func RetrieveToken(tokenString string, r *http.Request) (string, error) {
	tokenClaims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ServiceConfig(r).TokenKey), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("token is invalid")
	}
	expiresAt, ok := tokenClaims["exp"].(float64) // It was parsed to tokenClaims as float64
	if !ok {
		return "", errors.New("can't parse expiresAt")
	}
	if int64(expiresAt) < time.Now().Unix() { //Token is expired
		return "", errors.New("token is expired")
	}

	TokenUserAddress, ok := tokenClaims["address"].(string)
	if !ok {
		return "", errors.New("can't parse address")
	}
	return TokenUserAddress, nil
}

func getBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderSplit := strings.Split(authHeader, "Bearer ")
	if len(authHeaderSplit) != 2 {
		return "", errors.New("invalid Authorization header")
	}
	return authHeaderSplit[1], nil
}

func Authenticate(r *http.Request) (string, error) {
	sessToken, err := getBearerToken(r)
	if err != nil {
		err = errors.Wrap(err, "session token not found")
		return "", err
	}
	return sessToken, nil
}
