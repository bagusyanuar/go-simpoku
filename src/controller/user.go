package controller

import (
	"go-simpoku/database"
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func UserGet(c *gin.Context) {
	var user model.Member
	if err := database.DB.Debug().Where("id = ?", "9f2083d5-fc3e-469f-9feb-e3a599b83ed1").Omit("User").Preload("Specialist").First(&user).Error; err != nil {
		c.JSON(500, lib.Response{
			Code:    500,
			Data:    nil,
			Message: "failed " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    200,
		Data:    &user,
		Message: "success",
	})
}