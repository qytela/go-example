package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() *gorm.DB {
	var (
		host = os.Getenv("DATABASE_HOST")
		port = os.Getenv("DATABASE_PORT")
		user = os.Getenv("DATABASE_USER")
		pass = os.Getenv("DATABASE_PASS")
		name = os.Getenv("DATABASE_NAME")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	return db
}
