package configs

import (
	"fmt"
	"os"

	"github.com/PurinPintakhiew/Golang-API/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, databaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database:", err)
	} else {
		fmt.Println("connected to database")
		DB = db
		db.AutoMigrate(&models.Users{})
	}
}
