package main

import (
	"log"
	"os"

	"github.com/MartinAbelRR/blogbackend/database"
	"github.com/MartinAbelRR/blogbackend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files.")
	} else {
		log.Println("Connect successfully")
	}

	port := os.Getenv("PORT")

	app := fiber.New()

	routes.Setup(app)

	// Servidor
	app.Listen(":" + port)
}