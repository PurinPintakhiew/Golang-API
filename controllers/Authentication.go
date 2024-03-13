package controllers

import (
	"os"
	"time"

	"github.com/PurinPintakhiew/Golang-API/configs"
	"github.com/PurinPintakhiew/Golang-API/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// generate token
func GenerateToken(user models.Users) (string, error) {
	SecretKey := os.Getenv("SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id

	expirationTime := time.Now().Add(time.Hour * 24).Unix()
	claims["exp"] = expirationTime

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// register
func Register(c *fiber.Ctx) error {
	var data models.RegistrationData

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.Users
	existingUser := configs.DB.Where("email = ?", data.Email).First(&user)

	if existingUser.RowsAffected > 0 {
		return c.Status(409).JSON(fiber.Map{"status": false, "message": "Email already in use"})
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 15)

	newUser := models.Users{
		UserName: data.UserName,
		Email:    data.Email,
		Password: hashPassword,
	}

	configs.DB.Create(&newUser)

	token, err := GenerateToken(newUser)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status":      true,
		"users":       newUser,
		"accessToken": token,
	})

}

// login
func Login(c *fiber.Ctx) error {
	var data models.LoginData

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.Users
	if err := configs.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	token, err := GenerateToken(user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status":      true,
		"users":       user,
		"accessToken": token,
	})
}
