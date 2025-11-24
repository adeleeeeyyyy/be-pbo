package handlers

import (
	"be-pbo/config"
	"be-pbo/database"
	"be-pbo/dto"
	"be-pbo/mapper"
	"be-pbo/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var input dto.RegisterRequest

	// Parse body ke struct RegisterRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse json"})
	}

	// Validasi field wajib
	if input.Email == "" || input.Name == "" || input.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name, email, and password are required"})
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to hash password"})
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot create user"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "user created successfully",
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}


// Login
func Login(c *fiber.Ctx) error {
    var req dto.LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "cannot parse json"})
    }

    var user models.User
    if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "user not found"})
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "wrong password"})
    }

    token, err := config.GenerateJWT(user.ID, user.Email)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
			"error": "missing or invalid form jwt",
		})
    }

    return c.JSON(fiber.Map{
        "user": fiber.Map{
            "id":    user.ID,
            "name":  user.Name,
            "email": user.Email,
            "role":  user.Role,
        },
        "token": token,
    })
}

// Protected route
func Profile(c *fiber.Ctx) error {
	var user models.User
	userData := c.Locals("user")
	if userData == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userToken, ok := userData.(*jwt.Token)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims := userToken.Claims.(jwt.MapClaims)
	idVal, ok := claims["id"].(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID in token"})
	}
	UserID := uint(idVal)

	if err := database.DB.First(&user, UserID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	var response = mapper.UserToDTO(user)

	return c.Status(200).JSON(response)
}

