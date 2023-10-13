package dto

type CreateBookDTO struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
