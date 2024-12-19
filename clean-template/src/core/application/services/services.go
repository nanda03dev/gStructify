package services

import (
	"sync"

	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/repositories"
)

type Services struct {
	TemplateEntityService TemplateEntityService
}

var (
	once        sync.Once
	AllServices *Services
)

func GetServices() *Services {
	once.Do(func() {
		var AllRepositories = repositories.GetRepositories()
		AllServices = &Services{
			TemplateEntityService: NewTemplateEntityService(AllRepositories.TemplateEntityRepository),
		}
	})
	return AllServices
}
