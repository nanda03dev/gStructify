package bootstrap

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nanda03dev/go-ms-template/src/core/application/worker"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/core/interface/route"
)

type ApplicationManager interface {
	ConnectDatabase()
	DisconnectDatabase()
	Run()
}

type applicationManager struct {
	ctx      context.Context
	fiberApp *fiber.App
}

func NewApplicationManager(ctx context.Context, fiberApp *fiber.App) ApplicationManager {
	return &applicationManager{ctx: ctx, fiberApp: fiberApp}
}

func (app *applicationManager) ConnectDatabase() {
	db.ConnectAll()
}

func (app *applicationManager) DisconnectDatabase() {
	db.ConnectAll().DisconnectAll()
}

func (app *applicationManager) Run() {

	// Initialize workers
	log.Println("Starting worker...")
	worker.InitializeWorkers(app.ctx)

	// Initialize routes
	route.InitializeRoutes(app.fiberApp)
}
