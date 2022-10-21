package database

import (
	"jwt-react/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=123 dbname=go-jwt-auth port=5432 sslmode=disable TimeZone=Asia/Singapore"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to database!")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})

}
