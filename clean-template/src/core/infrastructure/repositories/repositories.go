package repositories

import (
	"sync"
)

type Repositories struct {
	EventRepository          EventRepository
	TemplateEntityRepository TemplateEntityRepository
}

var (
	repositoriesOnce sync.Once
	allRepositories  *Repositories
)

func GetRepositories() *Repositories {
	repositoriesOnce.Do(func() {
		allRepositories = &Repositories{
			EventRepository:          NewEventRepository(),
			TemplateEntityRepository: NewTemplateEntityRepository(),
		}
	})
	return allRepositories
}
