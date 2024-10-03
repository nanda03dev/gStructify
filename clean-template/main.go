package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nanda03dev/go-ms-template/src/app_module"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/interface/router"
)

func main() {
	var databases = db.ConnectAll()

	defer databases.DisconnectAll()

	app_module.InitModules()

	app := fiber.New()

	router.InitializeRoutes(app)

	app.Listen(":3000")

}
