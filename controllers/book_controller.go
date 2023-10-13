package controllers

import (
	"example/config"
	"example/dto"
	"example/helpers"
	"example/models"
	"example/resources"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct{}

func (b *BookController) Get(ctx *gin.Context) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	var bookResource []resources.BookResourceType
	for _, book := range books {
		bookResource = append(bookResource, resources.BookResource(book))
	}

	ctx.JSON(200, gin.H{
		"status": true,
		"data":   bookResource,
	})
}

func (b *BookController) Create(ctx *gin.Context) {
	var payload dto.CreateBookDTO

	var userId uint = ctx.GetUint("UserID")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	var book models.Book
	book.Title = payload.Title
	book.Author = payload.Author
	book.UserID = userId

	if err := config.DB.Create(&book).Error; err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status": true,
		"data":   resources.BookResource(book),
	})
}

func (b *BookController) Update(ctx *gin.Context) {
	var book models.Book
	var payload dto.UpdateBookDTO

	id, _ := strconv.Atoi(ctx.Param("id"))
	err := config.DB.First(&book, id).Error
	if err == gorm.ErrRecordNotFound {
		helpers.ErrorRecordNotFound(ctx)
		return
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	if payload.Title != "" {
		book.Title = payload.Title
	}
	if payload.Author != "" {
		book.Author = payload.Author
	}

	if err := config.DB.Save(&book).Error; err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status": true,
		"data":   resources.BookResource(book),
	})
}

func (b *BookController) Delete(ctx *gin.Context) {
	var book models.Book

	id, _ := strconv.Atoi(ctx.Param("id"))
	err := config.DB.First(&book, id).Error
	if err == gorm.ErrRecordNotFound {
		helpers.ErrorRecordNotFound(ctx)
		return
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		helpers.ErrorPanic(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status":  true,
		"message": "OK",
	})
}
