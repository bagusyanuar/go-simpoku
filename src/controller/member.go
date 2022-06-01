package controller

import (
	"go-simpoku/database"
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
	"go-simpoku/src/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type requestBody struct {
	Specialist []uint `json:"specialist"`
}

type response struct {
	ID        uuid.UUID  `json:"id"`
	UserID    uuid.UUID  `json:"user_id"`
	Name      string     `json:"name"`
	Phone     string     `json:"phone"`
	Avatar    string     `json:"avatar"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	User      model.User `json:"user"`
}

var member repository.Member

func Member(c *gin.Context) {
	_, e := lib.GetSignedUser(c)
	if e != nil {
		lib.Unauthorized(c)
	}
	data, err := member.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(500, lib.Response{
			Code:    500,
			Data:    nil,
			Message: "failed " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    200,
		Data:    data,
		Message: "success",
	})
}

func MemberProfile(c *gin.Context) {
	user_id, e := lib.GetSignedUser(c)
	if e != nil {
		lib.Unauthorized(c)
	}
	data, err := member.Find(user_id)
	if err != nil {
		c.AbortWithStatusJSON(500, lib.Response{
			Code:    500,
			Data:    nil,
			Message: "failed " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    200,
		Data:    data,
		Message: "success",
	})
}

func SetSpecialist(c *gin.Context) {
	var member model.Member
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
}
