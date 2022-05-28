package routes

import (
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
	}
	return route
}