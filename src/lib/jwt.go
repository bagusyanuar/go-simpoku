package lib

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var JWTSigninMethod = jwt.SigningMethodHS256
var JWTSignatureKey string = "ONLYGODKNOWS"
type JWT struct{}

type JWTClaims struct {
	jwt.StandardClaims
	Unique     uuid.UUID `json:"unique"`
	Identifier uuid.UUID `json:"identifier"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Role       string    `json:"roles"`
}

func (j JWT) GenerateToked(c JWTClaims) (string, error) {
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: "simpoku-app",
		},
		Unique: c.Unique,
		Username: c.Username,
		Email: c.Email,
		Role: c.Role,
		Identifier: c.Identifier,
	}

	token := jwt.NewWithClaims(JWTSigninMethod, claims)

	signedToken, err := token.SignedString([]byte(JWTSignatureKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}