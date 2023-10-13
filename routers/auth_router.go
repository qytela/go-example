package routers

import (
	"example/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.RouterGroup) {
	authController := new(controllers.AuthController)

	router.POST("/login", authController.Login)
	router.POST("/register", authController.Register)
	router.POST("/refresh", authController.RefreshToken)
}
