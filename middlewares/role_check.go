package middlewares

import (
	"example/config"
	"example/helpers"
	"example/models"

	"github.com/gin-gonic/gin"
)

func RoleCheck(requiredRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userRoles []models.UserHasRoles

		userId, _ := ctx.Get("UserID")
		if err := config.DB.Preload("Role").Where("user_id = ?", userId).Find(&userRoles).Error; err != nil {
			helpers.ErrorPanic(ctx, err)
			return
		}

		requiredRolesMap := make(map[string]bool)
		for _, role := range requiredRoles {
			requiredRolesMap[role] = true
		}

		var hasRequiredRole bool = false
		for _, userRole := range userRoles {
			if requiredRolesMap[userRole.Role.Name] {
				hasRequiredRole = true
				break
			}
		}

		if !hasRequiredRole {
			ctx.AbortWithStatusJSON(401, gin.H{
				"status":  false,
				"message": "User doesn't have any of the required roles",
			})
			return
		}

		ctx.Next()
	}
}
