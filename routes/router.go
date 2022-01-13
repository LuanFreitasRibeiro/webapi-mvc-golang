package routes

import (
	"github.com/LuanFreitasRibeiro/webapi-mvc-golang.git/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowAllBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}
