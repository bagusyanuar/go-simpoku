package routes

import (
	"go-simpoku/src/controller"

	"github.com/gin-gonic/gin"
)

func EventRoutes(route *gin.RouterGroup) {
	group := route.Group("/event")
	{
		group.GET("/", controller.IndexEvent)
		group.POST("/", controller.IndexEvent)
	}
}
