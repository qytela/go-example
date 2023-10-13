package helpers

import "github.com/gin-gonic/gin"

func ErrorPanic(ctx *gin.Context, err error) {
	if err != nil {
		ctx.JSON(422, gin.H{
			"status": false,
			"error":  err.Error(),
		})
	}
}

func ErrorEmailExists(ctx *gin.Context) {
	ctx.JSON(422, gin.H{
		"status":  false,
		"message": "Email already exists",
	})
}

func ErrorRecordNotFound(ctx *gin.Context) {
	ctx.JSON(404, gin.H{
		"status":  false,
		"message": "Record not found",
	})
}

func ErrorInvalidLogin(ctx *gin.Context) {
	ctx.JSON(403, gin.H{
		"status":  false,
		"message": "Invalid credentials",
	})
}
