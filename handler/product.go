package handler

import (
	"apisix-backend/service"
	"apisix-backend/share"
	"apisix-backend/structs"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var req structs.ProductGetRequest
	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	products, err := service.GetProductsService(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve products"})
	}
	return c.JSON(products)
}

func PostProduct(c *fiber.Ctx) error {
	var req structs.ProductPostRequest

	if err := share.ValidateAndBodyParser(c, &req); err != nil {
		return c.Status(err.(*fiber.Error).Code).JSON(fiber.Map{"error": err.Error()})
	}

	if err := service.CreateProductService(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
	}
	return c.JSON(fiber.Map{"success": true})
}

func PutProduct(c *fiber.Ctx) error {
	id, err := share.GetUintParam(c, "id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	var req structs.ProductPutRequest
	if err := share.ValidateAndBodyParser(c, &req); err != nil {
		return c.Status(err.(*fiber.Error).Code).JSON(fiber.Map{"error": err.Error()})
	}

	if err := service.UpdateProductService(id, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product"})
	}
	return c.JSON(fiber.Map{"success": true})
}

func PatchProduct(c *fiber.Ctx) error {
	id, err := share.GetUintParam(c, "id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	var req structs.ProductPatchRequest
	if err := share.ValidateAndBodyParser(c, &req); err != nil {
		return c.Status(err.(*fiber.Error).Code).JSON(fiber.Map{"error": err.Error()})
	}

	if err := service.PatchProductService(id, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to patch product"})
	}
	return c.JSON(fiber.Map{"success": true})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := share.GetUintParam(c, "id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	if err := service.DeleteProductService(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
	}
	return c.JSON(fiber.Map{"success": true})
}
