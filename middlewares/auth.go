package middlewares

import (
	"example/auth"
	"example/helpers"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("Authorization")
		if clientToken == "" {
			ctx.AbortWithStatusJSON(403, gin.H{
				"status":  false,
				"message": "No Authorization header provided",
			})
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			ctx.AbortWithStatusJSON(400, gin.H{
				"status":  false,
				"message": "Incorrect Format of Authorization Token",
			})
			return
		}

		jwtWrapper := auth.JwtWrapper{
			SecretKey: os.Getenv("JWT_SECRET_KEY"),
			Issuer:    os.Getenv("JWT_ISSUER"),
		}

		claims, err := jwtWrapper.ValidateToken(clientToken, jwtWrapper.SecretKey)
		if err != nil {
			helpers.ErrorPanic(ctx, err)
			ctx.Abort()
			return
		}

		ctx.Set("UserID", claims.UserID)
		ctx.Next()
	}
}
