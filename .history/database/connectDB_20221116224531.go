package database

import (
	"fmt"
	"github.com/PurinPintakhiew/Golang-API/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	} else {
		fmt.Println("connected database")
		DB = data
		data.AutoMigrate(&models.Cars{})
		data.AutoMigrate(&models.Users{})
	}
}
