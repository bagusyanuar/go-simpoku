package routes

import (
	"go-simpoku/src/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.RouterGroup) {
	group := route.Group("/auth")
	{
		admin := group.Group("/admin")
		{
			admin.POST("/sign-in", controller.AdminSignIn)
		}
	}
}