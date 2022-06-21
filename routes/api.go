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
		AuthController := controller.Auth{}
		MemberController := controller.Member{}
		EventController := controller.Event{}
		SpecialistController := controller.Specialist{}

		auth := api.Group("auth")
		{
			auth.POST("/sign-up", AuthController.SignUp)
			auth.POST("/sign-in", AuthController.SignIn)
		}

		member := api.Group("member")
		{
			member.GET("/", middleware.Auth, middleware.Member, MemberController.Index)
		}

		event := api.Group("event")
		{
			event.GET("/", EventController.Index)
			event.POST("/", EventController.Index)
			event.GET("/:slug", EventController.FindBySlug)
		}

		specialist := api.Group("specialist")
		{
			specialist.GET("/", SpecialistController.Index)
			specialist.POST("/", SpecialistController.Index)
			specialist.GET("/:id", SpecialistController.FindByID)
		}
	}
	return route
}
