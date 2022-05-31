package controller

import (
	"encoding/json"
	"go-simpoku/database"
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct{}

type userWithMember struct {
	model.BaseUser
	Member model.BaseMember `gorm:"foreignKey:UserID" json:"member"`
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

	user := model.BaseUser{
		Username: username,
		Email:    email,
		Password: &password,
		Roles:    roles,
	}

	member := model.BaseMember{
		Name: name,
	}

	userMember := userWithMember{
		BaseUser: user,
		Member:   member,
	}
	if err := tx.Create(&userMember).Error; err != nil {
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
		Unique:     userMember.ID,
		Identifier: userMember.Member.ID,
		Username:   userMember.Username,
		Email:      userMember.Email,
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
