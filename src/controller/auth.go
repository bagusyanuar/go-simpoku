package controller

import (
	"encoding/json"
	"go-simpoku/database"
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
	"go-simpoku/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct{}

type memberUser struct{
	model.Member
	User model.User
}

func SignUp(c *gin.Context) {
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")

	roles, _ := json.Marshal([]string{"member"})
	hash, errHashing := bcrypt.GenerateFromPassword([]byte(password), 13)
	if errHashing != nil {
		c.JSON(http.StatusInternalServerError, lib.Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Bad Request! Failed Hashing Password",
		})
		return
	}
	password = string(hash)

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user := model.User{
		Username: username,
		Email:    email,
		Password: &password,
		Roles:    roles,
	}

	member := memberUser{
		Member: model.Member{
			Name: name,
		},
		User: user,
	}

	if err := tx.Debug().Create(&member).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, lib.Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Error While Creating Account " + err.Error(),
		})
		return
	}

	jwt := lib.JWT{}
	claim := lib.JWTClaims{
		Unique:     member.UserID,
		Identifier: member.ID,
		// Username:   member.User.Username,
		// Email:      member.User.Email,
		Role:       "member",
	}
	accessToken, errorTokenize := jwt.GenerateToked(claim)
	if errorTokenize != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Error While Generate Token " + errorTokenize.Error(),
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, lib.Response{
		Code: http.StatusOK,
		Data: map[string]interface{}{
			"accessToken": accessToken,
		},
		Message: "Success Sign Up",
	})
}

func AdminSignIn(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	auth := repository.AdminAuth{
		User: model.User{
			Username: username,
			Password: &password,
		},
	}
	user, err := auth.SignIn()
	if err != nil {
		errorResponse := lib.ErrorSignIn(err)
		c.AbortWithStatusJSON(errorResponse.Code, errorResponse)
		return
	}
	jwt := lib.JWT{}
	claim := lib.JWTClaims{
		Unique:     user.User.ID,
		Identifier: user.Admin.ID,
		Username:   user.User.Username,
		Email:      user.User.Email,
		Role:       "admin",
	}
	accessToken, errorTokenize := jwt.GenerateToked(claim)
	if errorTokenize != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Error While Generate Token " + errorTokenize.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code: http.StatusOK,
		Data: map[string]interface{}{
			"accessToken": accessToken,
		},
		Message: "success sign in",
	})

}
