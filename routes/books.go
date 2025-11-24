package routes

import (
	"be-pbo/handlers"

	"github.com/gofiber/fiber/v2"
)

func BooksRoutes(router fiber.Router) {
	books := router.Group("/books")

	books.Post("/create", handlers.CreateBooks)
	books.Get("/show/:id", handlers.ShowBook)
	books.Get("/", handlers.GetBooks)
	books.Put("/edit/:id", handlers.UpdateBook)
}