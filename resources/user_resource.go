package resources

import (
	"example/models"
	"time"
)

type UserResourceType struct {
	ID        uint               `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Roles     []models.Role      `json:"roles"`
	Books     []BookResourceType `json:"books"`
}

func UserResource(user models.User) UserResourceType {
	var bookResource []BookResourceType
	for _, book := range user.Books {
		bookResource = append(bookResource, BookResource(book))
	}

	return UserResourceType{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Roles:     user.Roles,
		Books:     bookResource,
	}
}

type UserLoginResourceType struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func UserLoginResource(signedToken string, signedRefreshToken string) UserLoginResourceType {
	return UserLoginResourceType{
		Token:        signedToken,
		RefreshToken: signedRefreshToken,
	}
}
