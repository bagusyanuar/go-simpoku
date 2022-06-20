package routes

import (
	"go-simpoku/src/controller"
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
		// api.POST("/user/sign-up", controller.SignUp)
		// api.GET("/specialist", controller.Specialist)
		// api.POST("/specialist", controller.Specialist)
		// api.POST("/user/profile", controller.SetMemberProfile)
		// api.GET("/user/profile", controller.UserGet)
		// AuthRoutes(api)
		// MemberRoutes(api)
		
		routeAuth := api.Group("auth")
		{
			auth := controller.Auth{}
			routeAuth.POST("/sign-up", auth.SignIn)
			routeAuth.POST("/sign-in", auth.SignIn)
			// adminAuth := routeAuth.Group("admin")
			// {
			// 	adminAuth.POST("/sign-in", auth.AdminSignIn)
			// }

			// memberAuth := routeAuth.Group("member")
			// {
			// 	memberAuth.POST("/sign-up", auth.MemberSignUp)
			// 	memberAuth.POST("/sign-in", auth.SignIn)
			// }
		}

		routeEvent := api.Group("event")
		{
			event := controller.Event{}
			routeEvent.GET("/", event.Index)
			routeEvent.POST("/", event.Index)
		}
	}
	return route
}
