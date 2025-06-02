package route

import (
	"apisix-backend/handler"
	"apisix-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, prefix string) {
	public := app.Group(prefix)
	public.Post("/login", handler.Login)

	protected := app.Group(prefix, middleware.JWTProtected())
	protected.Get("/data", handler.GetProducts)
	protected.Post("/data", handler.PostProduct)
	protected.Put("/data/:id", handler.PutProduct)
	protected.Patch("/data/:id", handler.PatchProduct)
	protected.Delete("/data/:id", handler.DeleteProduct)
}
