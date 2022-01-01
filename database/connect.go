package database

import (
	"github.com/igor-koniukhov/fiber-go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect(){
	connection, err := gorm.Open(mysql.Open("test_igor:test_igor@/test_igor_db"), &gorm.Config{})
	if err !=nil {
		panic("Could not connect to the database!")
	}
	DB = connection
	connection.AutoMigrate(&models.User{}, models.PasswordReset{})
}