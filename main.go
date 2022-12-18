package main

import (
	"log"
	"os"

	"github.com/fprasty/GoApiWijaya/database"
	"github.com/fprasty/GoApiWijaya/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Can't load .env file")
	}
	port := os.Getenv("PORT")

	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
