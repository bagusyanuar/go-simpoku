package middleware

import (
	"go-simpoku/src/lib"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
		return
	}
	c.Set("user", claim)
	c.Next()
}

func Admin(c *gin.Context) {
	user := c.MustGet("user").(jwt.MapClaims)
	if user["roles"] != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, lib.Response{
			Code:    http.StatusForbidden,
			Data:    nil,
			Message: "Forbidden Access",
		})
		return
	}
	c.Next()
}

func Member(c *gin.Context) {
	user := c.MustGet("user").(jwt.MapClaims)
	if user["roles"] != "member" {
		c.AbortWithStatusJSON(http.StatusForbidden, lib.Response{
			Code:    http.StatusForbidden,
			Data:    nil,
			Message: "Forbidden Access",
		})
		return
	}
	c.Next()
}
