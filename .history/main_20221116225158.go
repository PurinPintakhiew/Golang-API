package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	
)

func main(){
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Post("search", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message":"Hello"})
	})

	app.Listen(":3001")
}