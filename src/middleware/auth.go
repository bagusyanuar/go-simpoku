package middleware

import (
	"go-simpoku/src/lib"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	jwt := lib.JWT{}
	claim, err := jwt.Claim(auth)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, lib.Response{
			Code:    http.StatusUnauthorized,
			Data:    nil,
			Message: strings.Title(err.Error()),
		})
	}
	c.Set("user", claim)
	c.Next()
}
