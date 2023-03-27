package provider

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type AuthCustomClaims struct {
	Auth           bool       `json:"auth"`
	Response       string     `json:"response"`
	jwt.StandardClaims
}

var signKey []byte = []byte("auth-server")

func SignToken(response string) (string, error) {
	var claims AuthCustomClaims = AuthCustomClaims {
		true,
		response,
		jwt.StandardClaims {
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(signKey)
}

func ParseTokenString(tokenString string) (string, error) {
	token, error := jwt.ParseWithClaims(tokenString, &AuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signKey, nil
	})

	if authCustomClaims, ok := token.Claims.(*AuthCustomClaims); ok && token.Valid {
		return authCustomClaims.Response, nil
	}
	return "", error
}
