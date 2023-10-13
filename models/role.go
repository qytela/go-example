package models

type Role struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}
