package utils

import "github.com/gofiber/fiber/v2"

type ResponseSuccess struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

type ResponseError struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Error   string `json:"error,omitempty"`
}

// Success response helper
func Success(c *fiber.Ctx, message string, data interface{}) error {
    return c.Status(fiber.StatusOK).JSON(ResponseSuccess{
        Success: true,
        Message: message,
        Data:    data,
    })
}

// Created response
func Created(c *fiber.Ctx, message string, data interface{}) error {
    return c.Status(fiber.StatusCreated).JSON(ResponseSuccess{
        Success: true,
        Message: message,
        Data:    data,
    })
}

// Error response helper
func Error(c *fiber.Ctx, status int, message string, err error) error {
    errMsg := ""
    if err != nil {
        errMsg = err.Error()
    }

    return c.Status(status).JSON(ResponseError{
        Success: false,
        Message: message,
        Error:   errMsg,
    })
}
