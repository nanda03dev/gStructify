package handlers

import "github.com/gofiber/fiber"

func SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{"data": data}
}

func ErrorResponse(data interface{}) fiber.Map {
	return fiber.Map{"error": data}
}
