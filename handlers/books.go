package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"

	"be-pbo/database"
	"be-pbo/dto"
	"be-pbo/helpers"
	"be-pbo/models"
	"be-pbo/utils"
)

func CreateBooks(c *fiber.Ctx) error {
    var input dto.BooksCreateRequest
    if err := c.BodyParser(&input); err != nil {
        return utils.Error(c, 400, "Invalid formdata", err)
    }

    // Upload Image
    imagePath, err := helpers.SaveUploadedFile(c, "image", "uploads/books")
    if err != nil {
        imagePath = ""
    }

    // Auto mapping pakai copier
    book := models.Books{}
    copier.Copy(&book, &input)

    // Tambahan manual
    book.Image = imagePath

    // Save
    if err := database.DB.Create(&book).Error; err != nil {
        return utils.Error(c, 500, "Failed to create book", err)
    }

    return utils.Created(c, "book created", book)
}

func ShowBook(c *fiber.Ctx) error {	
	id := c.Params("id")
	var book models.Books
	if err := database.DB.First(&book, id).Error; err != nil {
		return utils.Error(c, 404, "books not found", err)
	}

	return utils.Success(c, "books fetched", book)
}

func GetBooks(c *fiber.Ctx) error {
	var books []models.Books
	if err := database.DB.Find(&books).Error; err != nil {
		return utils.Error(c, 500, "failed to fetch books", err)
	}

	return utils.Success(c, "books fetched", books)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Books
	if err := database.DB.First(&book, id).Error; err != nil {
		return utils.Error(c, 500, "failed to fetch", err)
	}

	var input dto.BooksUpdateRequest
	if err := c.BodyParser(&input); err != nil {
		return utils.Error(c, 400, "invalid form data", err)
	}

	updates := utils.StructToMap(&input)

	newImage, err := helpers.UpdateUploadedFile(c, "image", "uploads/books", book.Image)
	if err != nil {
		return utils.Error(c, 500, "failed to save images", err)
	}

	updates["image"] = newImage

	if err := database.DB.Model(&book).Updates(updates).Error; err != nil {
		return utils.Error(c, 500, "failed to update book", err)
	}

	return utils.Success(c, "book updated", book)
}

