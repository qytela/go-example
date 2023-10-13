package routers

import (
	"example/controllers"
	"example/middlewares"

	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.RouterGroup) {
	bookController := new(controllers.BookController)

	bookRoutesProtected := router.Use(middlewares.Auth(), middlewares.RoleCheck([]string{"admin", "user"}))
	{
		bookRoutesProtected.GET("/", bookController.Get)
		bookRoutesProtected.POST("/", bookController.Create)
		bookRoutesProtected.PUT("/:id", bookController.Update)
		bookRoutesProtected.DELETE("/:id", bookController.Delete)
	}
}
