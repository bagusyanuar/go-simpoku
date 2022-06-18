package controller

import (
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
	"go-simpoku/src/repository"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)


func CreateEvent(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	start := c.PostForm("start")
	finish := c.PostForm("finish")

	startValue, _ := time.Parse("2006-01-02", start)
	finishValue, _ := time.Parse("2006-01-02", finish)
	event := repository.Event{
		Event: model.Event{
			Name: strings.Title(name),
			Slug: lib.MakeSlug(name),
			Description: description,
			StartAt: datatypes.Date(startValue),
			FinishAt: datatypes.Date(finishValue),
		},
	}
	data, err := event.Create()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Failed To insert Data " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: "success insert data",
	})
}