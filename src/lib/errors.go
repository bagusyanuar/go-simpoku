package lib

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func ErrorSignIn(err error) Response {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Response{
			Code:    http.StatusUnauthorized,
			Data:    nil,
			Message: "user not found!",
		}
	} else if errors.Is(err, ErrInvalidPassword) {
		return Response{
			Code:    http.StatusUnauthorized,
			Data:    nil,
			Message: "password did not match",
		}
	} else {
		return Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "error while sign in " + err.Error(),
		}
	}
}

func AbortInternalServerError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
		Code:    http.StatusInternalServerError,
		Data:    nil,
		Message: "internal server error : " + err.Error(),
	})
}
