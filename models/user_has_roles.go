package models

type UserHasRoles struct {
	UserID uint `json:"user_id" gorm:"foreignKey:id"`
	RoleID uint `json:"role_id" gorm:"foreignKey:id"`
	Role   Role `json:"role"`
}
