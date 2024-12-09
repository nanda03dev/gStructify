package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nanda03dev/go-ms-template/src/app_module"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/core/interface/router"
)

func main() {
	var databases = db.ConnectAll()

	defer databases.DisconnectAll()

	app_module.InitModules()

	app := fiber.New()

	router.InitializeRoutes(app)

	app.Listen(":3000")

}
