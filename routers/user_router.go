package routers

import (
	"example/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	userController := new(controllers.UserController)

	router.GET("/profile", userController.Profile)
	router.PUT("/", userController.Update)
}
