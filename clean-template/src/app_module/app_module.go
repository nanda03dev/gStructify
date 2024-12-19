package app_module

import (
	"context"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/nanda03dev/go-ms-template/src/core/application/workers"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/core/interface/routes"
)

var (
	once sync.Once
)

func ConnectDatabase() {
	db.ConnectAll()
}

func DisconnectDatabase() {
	db.ConnectAll().DisconnectAll()
}

func StartAppService(ctx context.Context, fiberApp *fiber.App) {

	// Initialize application modules
	log.Println("Starting workers...")
	workers.InitializeWorkers(ctx)

	routes.InitializeRoutes(fiberApp)
}
