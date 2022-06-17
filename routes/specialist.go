package routes

import (
	"go-simpoku/src/controller"
	"go-simpoku/src/middleware"

	"github.com/gin-gonic/gin"
)

func SpecialistRoute(route *gin.RouterGroup) {
	group := route.Group("/specialist")
	{
		group.GET("/", controller.Specialist)
		group.POST("/", middleware.Auth, middleware.Admin, controller.Specialist)
	}

}
