package controller

import (
	"encoding/json"
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
	"go-simpoku/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct{}

func (Auth) SignUp(c *gin.Context) {
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")

	roles, _ := json.Marshal([]string{"member"})
	hash, errHashing := bcrypt.GenerateFromPassword([]byte(password), 13)
	if errHashing != nil {
		lib.BadRequestError(c, errHashing)
		return
	}
	password = string(hash)

	auth := repository.AuthMember{
		User: model.User{
			Email: email,
			Username: username,
			Password: &password,
			Roles: roles,
		},
		Member: model.Member{
			Name: name,
		},
	}

	accessToken, err := auth.SignUp()
	if err != nil {
		lib.AbortInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code: http.StatusOK,
		Data: map[string]interface{}{
			"accessToken": accessToken,
		},
		Message: "Success Sign Up",
	})
}

func (Auth) SignIn(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	role := c.PostForm("role")
	auth := repository.Auth{
		User: model.User{
			Username: username,
			Password: &password,
		},
	}
	user, err := auth.SignIn(role)
	if err != nil {
		errorResponse := lib.ErrorSignIn(err)
		c.AbortWithStatusJSON(errorResponse.Code, errorResponse)
		return
	}
	jwt := lib.JWT{}
	claim := lib.JWTClaims{
		Unique:     user.User.ID,
		Identifier: user.Member.ID,
		Username:   user.User.Username,
		Email:      user.User.Email,
		Role:       "admin",
	}
	accessToken, errorTokenize := jwt.GenerateToked(claim)
	if errorTokenize != nil {
		lib.AbortInternalServerError(c, errorTokenize)
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code: http.StatusOK,
		Data: map[string]interface{}{
			"accessToken": accessToken,
			"user": user,
		},
		Message: "success sign in",
	})
}

// func (Auth) AdminSignIn(c *gin.Context) {
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")
// 	auth := repository.Auth{
// 		User: model.User{
// 			Username: username,
// 			Password: &password,
// 		},
// 	}
// 	user, err := auth.SignIn("admin")
// 	if err != nil {
// 		errorResponse := lib.ErrorSignIn(err)
// 		c.AbortWithStatusJSON(errorResponse.Code, errorResponse)
// 		return
// 	}
// 	jwt := lib.JWT{}
// 	claim := lib.JWTClaims{
// 		Unique:     user.User.ID,
// 		Identifier: user.Admin.ID,
// 		Username:   user.User.Username,
// 		Email:      user.User.Email,
// 		Role:       "admin",
// 	}
// 	accessToken, errorTokenize := jwt.GenerateToked(claim)
// 	if errorTokenize != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"code":    http.StatusInternalServerError,
// 			"data":    nil,
// 			"message": "Error While Generate Token " + errorTokenize.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, lib.Response{
// 		Code: http.StatusOK,
// 		Data: map[string]interface{}{
// 			"accessToken": accessToken,
// 		},
// 		Message: "success sign in",
// 	})

// }
