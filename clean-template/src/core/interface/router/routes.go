package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nanda03dev/go-ms-template/src/app_module"
)

func InitializeRoutes(app *fiber.App) {

	var appModule = app_module.GetAppModule()

	templateEntityHandler := appModule.Handler.TemplateEntityHandler
	api := app.Group("/api")
	templateEntityV1Routes := api.Group("/v1/templateEntity")
	templateEntityV1Routes.Post("/", templateEntityHandler.CreateTemplateEntity)
	templateEntityV1Routes.Get("/:id", templateEntityHandler.GetTemplateEntityByID)
	templateEntityV1Routes.Put("/:id", templateEntityHandler.UpdateTemplateEntityById)
	templateEntityV1Routes.Delete("/:id", templateEntityHandler.DeleteTemplateEntityById)

}
