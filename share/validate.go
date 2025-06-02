package share

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateAndBodyParser(c *fiber.Ctx, req interface{}) error {
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := validate.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}

	return nil
}

func GetUintParam(c *fiber.Ctx, name string) (uint, error) {
	idStr := c.Params(name)
	id, err := strconv.ParseUint(idStr, 10, 64)
	return uint(id), err
}
