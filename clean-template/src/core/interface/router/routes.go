package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nanda03dev/go-ms-template/src/app_module"
)

func InitializeRoutes(app *fiber.App) {

	var appModule = app_module.GetAppModule()
	api := app.Group("/api")

	// TemplateEntity CRUD API'S
	templateEntityHandler := appModule.Handler.TemplateEntityHandler
	templateEntityV1Routes := api.Group("/v1/templateEntity")
	templateEntityV1Routes.Post("/", templateEntityHandler.CreateTemplateEntity)
	templateEntityV1Routes.Post("/filter", templateEntityHandler.FindTemplateEntityWithFilter)
	templateEntityV1Routes.Get("/:id", templateEntityHandler.GetTemplateEntityByID)
	templateEntityV1Routes.Put("/:id", templateEntityHandler.UpdateTemplateEntityById)
	templateEntityV1Routes.Delete("/:id", templateEntityHandler.DeleteTemplateEntityById)

}
