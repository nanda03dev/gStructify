package app_module

import (
	"context"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/nanda03dev/go-ms-template/src/core/application/services"
	"github.com/nanda03dev/go-ms-template/src/core/application/workers"
	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/repositories"
	"github.com/nanda03dev/go-ms-template/src/core/interface/handlers"
)

type Repository struct {
	TemplateEntityRepository aggregates.TemplateEntityRepository
}

type Service struct {
	TemplateEntityService services.TemplateEntityService
}

type Handler struct {
	TemplateEntityHandler handlers.TemplateEntityHandler
}

type Module struct {
	Service    Service
	Repository Repository
	Handler    Handler
}

var (
	once    sync.Once
	modules *Module
)

func InitializeModules() *Module {
	return GetModule()
}

func GetModule() *Module {
	once.Do(func() {

		var AllRepositories = Repository{
			TemplateEntityRepository: repositories.NewTemplateEntityRepository(),
		}

		var AllServices = Service{
			TemplateEntityService: services.NewTemplateEntityService(AllRepositories.TemplateEntityRepository),
		}

		var AllHandlers = Handler{
			TemplateEntityHandler: handlers.NewTemplateEntityHandler(AllServices.TemplateEntityService),
		}

		modules = &Module{
			Repository: AllRepositories,
			Service:    AllServices,
			Handler:    AllHandlers,
		}
	})
	return modules
}

func ConnectDatabase() {
	db.ConnectAll()
}

func DisconnectDatabase() {
	db.ConnectAll().DisconnectAll()
}

func StartAppService(ctx context.Context, fiberApp *fiber.App) {
	InitializeModules()

	// Initialize application modules
	log.Println("Starting workers...")
	workers.InitializeWorkers(ctx)

	InitializeRoutes(fiberApp)
}
