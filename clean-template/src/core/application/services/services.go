package services

import (
	"sync"
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
		var AllRepositories = repositories.GetRepositories()
		allServices = &Services{
			TemplateEntityService: NewTemplateEntityService(AllRepositories.TemplateEntityRepository),
			EventService:          NewEventService(AllRepositories.EventRepository),
		}
	})
	return allServices
}
