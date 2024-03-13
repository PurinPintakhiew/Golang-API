package main

import (
	"github.com/PurinPintakhiew/Golang-API/configs"
	"github.com/PurinPintakhiew/Golang-API/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	configs.ConnectDB()
	routes.Route()
}
