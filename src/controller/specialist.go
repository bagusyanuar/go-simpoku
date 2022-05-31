package controller

import (
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
	"go-simpoku/src/repository"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Specialist(c *gin.Context) {
	specialist := repository.Specialist{}

	if c.Request.Method == "POST" {
		name := c.PostForm("name")

		formData := model.BaseSpecialist{
			Name: strings.Title(name),
			Slug: lib.MakeSlug(name),
		}

		data, err := specialist.Create(&formData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, lib.Response{
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
		return
	}
	result, err := specialist.FindAll()
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
		Message: "success fetch data",
	})
}
