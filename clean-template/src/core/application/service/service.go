package service

import (
	"sync"

	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/repository"
)

type Services struct {
	TemplateEntityService TemplateEntityService
	EventService          EventService
}

var (
	servicesOnce sync.Once
	allServices  *Services
)

func GetServices() *Services {
	servicesOnce.Do(func() {
		var AllRepository = repository.GetRepositories()
		allServices = &Services{
			TemplateEntityService: NewTemplateEntityService(AllRepository.TemplateEntityRepository),
			EventService:          NewEventService(AllRepository.EventRepository),
		}
	})
	return allServices
}
