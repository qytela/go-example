package controllers

import (
	"example/config"
	"example/dto"
	"example/helpers"
	"example/models"
	"example/resources"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct{}

func (u *UserController) Profile(ctx *gin.Context) {
	var user models.User

	userId, _ := ctx.Get("UserID")
	err := config.DB.Preload("Roles").Preload("Books").First(&user, userId).Error
	if err == gorm.ErrRecordNotFound {
		helpers.ErrorRecordNotFound(ctx)
		return
	}

	ctx.JSON(200, gin.H{
		"status": true,
		"data":   resources.UserResource(user),
	})
}

func (u *UserController) Update(ctx *gin.Context) {
	var user models.User
	var payload dto.UpdateUserDTO

	userId, _ := ctx.Get("UserID")
	err := config.DB.First(&user, userId).Error
	if err == gorm.ErrRecordNotFound {
		helpers.ErrorRecordNotFound(ctx)
		return
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(500, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	if payload.Name != "" {
		user.Name = payload.Name
	}
	if payload.Email != "" {
		user.Email = payload.Email
	}
	if payload.Password != "" {
		user.Password = payload.Password
	}

	if err := config.DB.Save(&user).Error; err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status": true,
		"data":   resources.UserResource(user),
	})
}
