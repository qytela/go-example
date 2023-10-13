package utils

import (
	"example/auth"
	"example/config"
	"example/models"
	"os"
	"strconv"

	"gorm.io/gorm"
)

func GenerateToken(userId uint) (signedToken string, signedRefreshToken string, err error) {
	expirationMinutes, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRATION_MINUTES"), 10, 64)
	expirationHours, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRATION_HOURS"), 10, 64)
	jwtWrapper := auth.JwtWrapper{
		SecretKey:         os.Getenv("JWT_SECRET_KEY"),
		SecretKeyRefresh:  os.Getenv("JWT_SECRET_KEY_REFRESH"),
		Issuer:            os.Getenv("JWT_ISSUER"),
		ExpirationMinutes: expirationMinutes,
		ExpirationHours:   expirationHours,
	}

	signedToken, err = jwtWrapper.GenerateToken(userId)
	if err != nil {
		return
	}

	signedRefreshToken, err = jwtWrapper.GenerateRefreshToken(userId)
	if err != nil {
		return
	}

	return signedToken, signedRefreshToken, err
}

func ValidateRefreshToken(clientToken string) (signedToken string, err error) {
	expirationMinutes, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRATION_MINUTES"), 10, 64)
	expirationHours, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRATION_HOURS"), 10, 64)
	jwtWrapper := auth.JwtWrapper{
		SecretKey:         os.Getenv("JWT_SECRET_KEY"),
		SecretKeyRefresh:  os.Getenv("JWT_SECRET_KEY_REFRESH"),
		Issuer:            os.Getenv("JWT_ISSUER"),
		ExpirationMinutes: expirationMinutes,
		ExpirationHours:   expirationHours,
	}

	claims, err := jwtWrapper.ValidateToken(clientToken, jwtWrapper.SecretKeyRefresh)
	if err != nil {
		return
	}

	signedToken, err = jwtWrapper.GenerateToken(claims.UserID)
	if err != nil {
		return
	}

	return signedToken, err
}

func UserEmailExists(email string, user *models.User) bool {
	result := config.DB.Where("email = ?", email).First(&user)
	return result.Error != gorm.ErrRecordNotFound
}

func AssignDefaultUserRole(user *models.User) error {
	var defaultRole models.Role

	if err := config.DB.Where("name = ?", "user").First(&defaultRole).Error; err != nil {
		return err
	}

	userRoles := models.UserHasRoles{
		UserID: user.ID,
		RoleID: defaultRole.ID,
	}

	if err := config.DB.Create(&userRoles).Error; err != nil {
		return err
	}

	return nil
}
