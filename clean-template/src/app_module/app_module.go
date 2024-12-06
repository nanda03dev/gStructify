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

		var AllRepositories = Repository{
			TemplateEntityRepository: repositories.NewTemplateEntityRepository(databases),
		}

		var AllServices = Service{
			TemplateEntityService: services.NewTemplateEntityService(AllRepositories.TemplateEntityRepository),
		}

		var AllHandlers = Handler{
			TemplateEntityHandler: handlers.NewTemplateEntityHandler(AllServices.TemplateEntityService),
		}

		appModule = &AppModule{
			Repository: AllRepositories,
			Service:    AllServices,
			Handler:    AllHandlers,
		}
	})
	return appModule
}
