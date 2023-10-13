package controllers

import (
	"example/config"
	"example/dto"
	"example/helpers"
	"example/models"
	"example/resources"
	"example/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (u *AuthController) Login(ctx *gin.Context) {
	var user models.User
	var payload dto.LoginUserDTO

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	if isExists := utils.UserEmailExists(payload.Email, &user); !isExists {
		helpers.ErrorInvalidLogin(ctx)
		return
	}

	if err := user.ComparePassword(payload.Password); err != nil {
		helpers.ErrorInvalidLogin(ctx)
		return
	}

	signedToken, signedRefreshToken, err := utils.GenerateToken(user.ID)
	if err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status": true,
		"data":   resources.UserLoginResource(signedToken, signedRefreshToken),
	})
}

func (u *AuthController) Register(ctx *gin.Context) {
	var payload dto.CreateUserDTO

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	var existingUser models.User
	if isExists := utils.UserEmailExists(payload.Email, &existingUser); isExists {
		helpers.ErrorEmailExists(ctx)
		return
	}

	var user models.User
	user.Name = payload.Name
	user.Email = payload.Email
	user.Password = payload.Password

	if err := config.DB.Create(&user).Error; err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	userRoles := models.UserHasRoles{
		UserID: user.ID,
		RoleID: 1,
	}

	if err := config.DB.Create(&userRoles).Error; err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status": true,
		"data":   resources.UserResource(user),
	})
}

func (u *AuthController) RefreshToken(ctx *gin.Context) {
	var payload dto.RefreshTokenDTO

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	signedToken, err := utils.ValidateRefreshToken(payload.RefreshToken)
	if err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status": true,
		"data":   resources.UserLoginResource(signedToken, ""),
	})
}
