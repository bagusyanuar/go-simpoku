package lib

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrBearerType      = errors.New("invalid bearer type")
	ErrSignInMethod    = errors.New("invalid signin method")
	ErrJWTClaims       = errors.New("invalid jwt claim")
	ErrJWTParse        = errors.New("invalid parse jwt")
	ErrNoAuthorization = errors.New("unauthorized")
	ErrInvalidPassword = errors.New("password did not match")
)

func Unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, Response{
		Code:    http.StatusUnauthorized,
		Data:    nil,
		Message: "Unautthorized : failed to get signed user",
	})
	return
}
