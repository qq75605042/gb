package routers

import (
	"blog/api/controllers"

	"github.com/gin-gonic/gin"
)

func LoadApiRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"version": "1.0",
			})
		})
		api.GET("/articles", controllers.GetArticles)
	}
}
