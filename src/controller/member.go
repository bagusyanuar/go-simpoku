package controller

import (
	"go-simpoku/database"
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type requestBody struct {
	Specialist []uint `json:"specialist"`
}

func MemberProfile(c *gin.Context) {
	user := c.MustGet("user").(jwt.MapClaims)
	var member model.Member
	user_id, e := uuid.Parse(user["identifier"].(string))
	if e != nil {
		c.JSON(http.StatusUnauthorized, lib.Response{
			Code:    http.StatusUnauthorized,
			Data:    nil,
			Message: "Internal Server Error. Failed To Get User",
		})
		return
	}
	if err := database.DB.Debug().Preload("Specialist").Preload("User").Where("id = ?", user_id).First(&member).Error; err != nil {
		c.JSON(500, lib.Response{
			Code:    500,
			Data:    nil,
			Message: "failed " + err.Error(),
		})
		return
	}

	if c.Request.Method == "POST" {
		var request requestBody
		c.BindJSON(&request)

		specialists := []model.Specialist{}

		for _, v := range request.Specialist {
			specialists = append(specialists, model.Specialist{
				ID: v,
			})
		}
		err := database.DB.Debug().Model(&member).Association("Specialist").Append(specialists)

		if err != nil {
			c.JSON(http.StatusInternalServerError, lib.Response{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: "failed to patch specialist " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, lib.Response{
			Code:    200,
			Data:    nil,
			Message: "success",
		})
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    200,
		Data:    &member,
		Message: "success",
	})
}
