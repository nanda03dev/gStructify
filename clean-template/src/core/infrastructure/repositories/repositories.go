package repositories

import (
	"sync"

	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregates"
)

type Repositories struct {
	TemplateEntityRepository aggregates.TemplateEntityRepository
}

var (
	once            sync.Once
	AllRepositories *Repositories
)

func GetRepositories() *Repositories {
	once.Do(func() {
		AllRepositories = &Repositories{
			TemplateEntityRepository: NewTemplateEntityRepository(),
		}
	})
	return AllRepositories
}
