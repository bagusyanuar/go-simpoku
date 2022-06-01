package routes

import (
	"go-simpoku/src/controller"
	"go-simpoku/src/middleware"

	"github.com/gin-gonic/gin"
)

func MemberRoutes(route *gin.RouterGroup) {
	group := route.Group("/members")
	{
		group.GET("/", controller.Member)
		group.GET("/me", middleware.Auth, middleware.Member, controller.MemberProfile)
	}
}
