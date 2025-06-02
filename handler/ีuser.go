package handler

import (
	"apisix-backend/service"
	"apisix-backend/share"
	"apisix-backend/structs"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req structs.LoginRequest

	if err := share.ValidateAndBodyParser(c, &req); err != nil {
		return c.Status(err.(*fiber.Error).Code).JSON(fiber.Map{"error": err.Error()})
	}

	token, user, err := service.LoginService(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	return c.JSON(fiber.Map{"success": true, "token": token, "id": user.ID, "username": user.Username})
}
