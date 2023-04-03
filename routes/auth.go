package routes

import (
	"github.com/PurinPintakhiew/Golang-API/controllers"
	"github.com/PurinPintakhiew/Golang-API/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Route() {
	database.ConnectDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// public display
	app.Static("/static", "./public")

	// route path Hello
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, Purin"})
	})

	// Authentication api ประกาศ
	app.Post("/register", controllers.Register) // ลงทะเบียนหรือเพิ่มผู้ใช้
	app.Post("/login", controllers.Login)       // ล็อคอิน
	app.Post("/logout", controllers.Logout)     // ออกจากระบบ

	app.Listen((":8080"))
}
