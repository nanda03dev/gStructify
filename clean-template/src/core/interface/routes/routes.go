package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nanda03dev/go-ms-template/src/core/interface/handlers"
	"github.com/nanda03dev/go-ms-template/src/core/interface/middlewares"
)

func InitializeRoutes(fiberApp *fiber.App) {
	// Apply the global recovery middleware first
	fiberApp.Use(middlewares.RecoveryMiddleware())
	fiberApp.Use(healthcheck.New())
	fiberApp.Use(logger.New())

	api := fiberApp.Group("/api")

	AllHandlers := handlers.GetHandlers()

	// TemplateEntity CRUD API'S
	templateEntityHandler := AllHandlers.TemplateEntityHandler
	templateEntityV1Routes := api.Group("/v1/templateEntity")
	templateEntityV1Routes.Post("/", templateEntityHandler.CreateTemplateEntity)
	templateEntityV1Routes.Post("/filter", templateEntityHandler.FindTemplateEntityWithFilter)
	templateEntityV1Routes.Get("/:id", templateEntityHandler.GetTemplateEntityByID)
	templateEntityV1Routes.Put("/:id", templateEntityHandler.UpdateTemplateEntityById)
	templateEntityV1Routes.Delete("/:id", templateEntityHandler.DeleteTemplateEntityById)

}
