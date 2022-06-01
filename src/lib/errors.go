package lib

import "errors"

var (
	ErrBearerType      = errors.New("invalid bearer type")
	ErrSignInMethod    = errors.New("invalid signin method")
	ErrJWTClaims       = errors.New("invalid jwt claim")
	ErrJWTParse        = errors.New("invalid parse jwt")
	ErrNoAuthorization = errors.New("unauthorized")
	ErrInvalidPassword = errors.New("password did not match")
)
