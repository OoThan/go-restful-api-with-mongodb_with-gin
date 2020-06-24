package routes

import (
	"github.com/OoThan/go-restful-api-with-mongodb/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartGin() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/books", controllers.GetAllBooks)
		api.POST("/books", controllers.CreateBook)
		api.GET("/books/:id", controllers.GetBook)
		api.PUT("/books/:id", controllers.UpdateBook)
		api.DELETE("/books/:id", controllers.DeleteBook)
	}
	router.NoRoute(func(context *gin.Context) {
		context.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8088")
}
