package controller

import (
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
	"go-simpoku/src/repository"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Event struct{}

func IndexEvent(c *gin.Context) {

	if c.Request.Method == "POST" {

		file, _ := c.FormFile("file")
		if file == nil {
			c.JSON(http.StatusOK, lib.Response{
				Code:    http.StatusOK,
				Data:    nil,
				Message: "success",
			})
			return
		}
		ext := filepath.Ext(file.Filename)
		fileName := "assets/icons/" + uuid.New().String() + ext
		if errUpload := c.SaveUploadedFile(file, fileName); errUpload != nil {
			c.JSON(http.StatusInternalServerError, lib.Response{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: "Internal Server Error. Failed while upload icon" + errUpload.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, lib.Response{
			Code: http.StatusOK,
			Data: map[string]interface{}{
				"file": file,
				"name": file.Filename,
				"ext":  ext,
			},
			Message: "success",
		})
		return
	}
	event := repository.Event{}
	data, err := event.FindAll()
	if err != nil {
		lib.AbortInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: "success",
	})
}

func CreateEvent(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	start := c.PostForm("start")
	finish := c.PostForm("finish")

	startValue, _ := time.Parse("2006-01-02", start)
	finishValue, _ := time.Parse("2006-01-02", finish)
	event := repository.Event{
		Event: model.Event{
			Name:        strings.Title(name),
			Description: description,
			Slug:        lib.MakeSlug(name),
			Image:       "",
			StartAt:     datatypes.Date(startValue),
			FinishAt:    datatypes.Date(finishValue),
			Location:    "",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
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
