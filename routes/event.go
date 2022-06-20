package routes

import (
	"go-simpoku/src/controller"

	"github.com/gin-gonic/gin"
)

func EventRoutes(route *gin.RouterGroup) {
	event := controller.Event{}
	group := route.Group("/event")
	{
		group.GET("/", event.Index)
		group.POST("/", event.Index)
	}
}
