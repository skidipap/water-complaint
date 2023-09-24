package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("secret123")

type JWTClaim struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

func ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) { return jwtSecret, nil })
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err := errors.New("could'nt parse claim")
		return err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err := errors.New("token expired")
		return err
	}
	return nil
}
