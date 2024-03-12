package main

import (
	"github.com/PurinPintakhiew/Golang-API/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	routes.Route()
}
