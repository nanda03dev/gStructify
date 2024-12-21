package repositories

import (
	"sync"
)

type Repositories struct {
	EventRepository          EventRepository
	TemplateEntityRepository TemplateEntityRepository
}

var (
	once            sync.Once
	AllRepositories *Repositories
)

func GetRepositories() *Repositories {
	once.Do(func() {
		AllRepositories = &Repositories{
			EventRepository:          NewEventRepository(),
			TemplateEntityRepository: NewTemplateEntityRepository(),
		}
	})
	return AllRepositories
}
