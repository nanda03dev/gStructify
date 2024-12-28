package repository

import (
	"sync"

	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/db"
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
	var databases = db.ConnectAll()
	repositoriesOnce.Do(func() {
		allRepositories = &Repositories{
			EventRepository:          NewEventRepository(databases),
			TemplateEntityRepository: NewTemplateEntityRepository(databases),
		}
	})
	return allRepositories
}
