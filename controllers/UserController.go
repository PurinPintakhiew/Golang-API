package controllers

import (
	"github.com/PurinPintakhiew/Golang-API/configs"
	"github.com/PurinPintakhiew/Golang-API/models"
	"github.com/gofiber/fiber/v2"
)

func Users(c *fiber.Ctx) error {
	var users []models.Users

	result := configs.DB.Find(&users)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "Internal Server Error"})
	}

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Data not found"})
	}

	return c.JSON(fiber.Map{"status": true, "data": users})
}
