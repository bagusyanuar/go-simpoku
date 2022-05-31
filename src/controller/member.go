package controller

import (
	"go-simpoku/database"
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestBody struct {
	ID         string `json:"id"`
	Specialist []uint `json:"specialist"`
}

func SetMemberProfile(c *gin.Context) {
	var request requestBody
	c.BindJSON(&request)

	var member model.BaseMember
	if err := database.DB.Debug().Where("id = ?", request.ID).First(&member).Error; err != nil {
		c.JSON(500, lib.Response{
			Code:    500,
			Data:    nil,
			Message: "failed " + err.Error(),
		})
		return
	}

	database.DB.Debug().Model(&member).Association("Specialist").Append([]model.BaseSpecialist{
		{ID: uint(request.Specialist[0])},
		{ID: uint(request.Specialist[1])},
	})

	// data := model.Member{
	// 	BaseMember: member,
	// 	Specialist: []model.BaseSpecialist{
	// 		{ID: uint(request.Specialist[0])},
	// 		{ID: uint(request.Specialist[1])},
	// 	},
	// }

	if err := database.DB.Debug().Save(&member).Error; err != nil {
		c.JSON(500, lib.Response{
			Code:    500,
			Data:    nil,
			Message: "failed " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    200,
		Data:    &request,
		Message: "success",
	})
}
