package app_module

import (
	"sync"
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

type AppModule struct {
	Service    Service
	Repository Repository
	Handler    Handler
}

var (
	once      sync.Once
	appModule *AppModule
)

func InitModules() *AppModule {
	return GetAppModule()
}

func GetAppModule() *AppModule {
	once.Do(func() {
		var databases = db.ConnectAll()

		templateEntityRepository := repositories.NewTemplateEntityRepository(databases)
		templateEntityService := services.NewTemplateEntityService(templateEntityRepository)
		templateEntityHandler := handlers.NewTemplateEntityHandler(templateEntityService)

		appModule = &AppModule{
			Repository: Repository{
				TemplateEntityRepository: templateEntityRepository,
			},
			Service: Service{
				TemplateEntityService: templateEntityService,
			},
			Handler: Handler{
				TemplateEntityHandler: templateEntityHandler,
			},
		}
	})
	return appModule
}
