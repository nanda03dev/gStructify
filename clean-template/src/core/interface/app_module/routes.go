package app_module

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nanda03dev/go-ms-template/src/core/interface/handlers"
)

func InitializeRoutes(fiberApp *fiber.App) {
	modules := GetModule()
	api := fiberApp.Group("/api")

	// TemplateEntity CRUD API'S
	templateEntityHandler := modules.Handler.TemplateEntityHandler
	templateEntityV1Routes := api.Group("/v1/templateEntity")
	templateEntityV1Routes.Post("/", templateEntityHandler.CreateTemplateEntity)
	templateEntityV1Routes.Post("/filter", templateEntityHandler.FindTemplateEntityWithFilter)
	templateEntityV1Routes.Get("/:id", templateEntityHandler.GetTemplateEntityByID)
	templateEntityV1Routes.Put("/:id", templateEntityHandler.UpdateTemplateEntityById)
	templateEntityV1Routes.Delete("/:id", templateEntityHandler.DeleteTemplateEntityById)

}
