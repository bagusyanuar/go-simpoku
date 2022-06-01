package routes

import (
	"go-simpoku/src/controller"
	"go-simpoku/src/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	route := gin.Default()
	route.SetTrustedProxies([]string{"localhost"})
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("assets"))))
	api := route.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(202, gin.H{
				"code": "OK",
			})
		})
		api.POST("/user/sign-up", controller.SignUp)
		api.GET("/specialist", controller.Specialist)
		api.POST("/specialist", controller.Specialist)
		api.GET("/user/profile", middleware.Auth, controller.MemberProfile)
		api.POST("/user/profile", middleware.Auth, controller.MemberProfile)
	}
	return route
}
