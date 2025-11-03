package main

import (
	"be-pbo/database"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}

	database.Connect()

	// database.DB.Automigrate()

	app := fiber.New()

	app.Static("/uploads", "./uploads")

	
}