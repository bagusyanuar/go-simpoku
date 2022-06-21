package controller

import (
	"errors"
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
	"go-simpoku/src/repository"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Specialist struct{
	model.Specialist
}

func (Specialist) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		name := c.PostForm("name")
		specialist := repository.Specialist{
			Specialist: model.Specialist{
				Name: strings.Title(name),
				Slug: lib.MakeSlug(name),
			},
		}
		data, err := specialist.Create()
		if err != nil {
			c.JSON(http.StatusInternalServerError, lib.Response{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: "failed to insert data : " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, lib.Response{
			Code:    http.StatusOK,
			Data:    data,
			Message: "success",
		})
		return
	}
	specialist := repository.Specialist{}
	q := c.Query("q")
	result, err := specialist.FindAll(q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Failed To Fetch Data " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Data:    result,
		Message: "success",
	})
}

func (Specialist) FindByID(c *gin.Context)  {
	id := c.Param("id")
	specialist := repository.Specialist{}
	data, err := specialist.Find(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			lib.RecordNotFound(c)
			return
		}
		lib.AbortInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: "success",
	})
}
