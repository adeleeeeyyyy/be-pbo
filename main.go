package main

import (
	"be-pbo/database"
	"be-pbo/models"
	"be-pbo/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}

	database.Connect()

	database.DB.AutoMigrate(
		&models.User{},
		&models.Books{}, 
	)

	app := fiber.New()

	app.Static("/uploads", "./uploads")

	routes.SetupRoutes(app)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("server running on %s\n", addr)

	if err := app.Listen(addr); err != nil {
		log.Fatalf("failed to start fiber: %v", err)
	}
}