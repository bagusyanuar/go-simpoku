package controller

import (
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
	"go-simpoku/src/repository"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Event struct{}

func (Event) Index(c *gin.Context) {

	if c.Request.Method == "POST" {

		name := c.PostForm("name")
		description := c.PostForm("description")
		start := c.PostForm("start")
		finish := c.PostForm("finish")
		location := c.PostForm("location")

		startValue, _ := time.Parse("2006-01-02", start)
		finishValue, _ := time.Parse("2006-01-02", finish)

		file, _ := c.FormFile("image")
		if file == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
				Code:    http.StatusBadRequest,
				Data:    nil,
				Message: "file image required",
			})
			return
		}
		ext := filepath.Ext(file.Filename)
		fileName := "assets/images/" + uuid.New().String() + ext
		if errUpload := c.SaveUploadedFile(file, fileName); errUpload != nil {
			if _, err := os.Stat("/assets/images"); err != nil {
				lib.AbortInternalServerError(c, err)
				return
			}
			lib.AbortInternalServerError(c, errUpload)
			return
		}
		event := repository.Event{
			Event: model.Event{
				Name:        strings.Title(name),
				Description: description,
				Slug:        lib.MakeSlug(name),
				Image:       fileName,
				StartAt:     datatypes.Date(startValue),
				FinishAt:    datatypes.Date(finishValue),
				Location:    location,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
		}
		data, err := event.Create()
		if err != nil {
			lib.AbortInternalServerError(c, err)
			return
		}
		c.JSON(http.StatusOK, lib.Response{
			Code:    http.StatusOK,
			Data:    data,
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