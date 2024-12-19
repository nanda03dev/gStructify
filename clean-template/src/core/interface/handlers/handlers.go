package handlers

import (
	"sync"

	"github.com/nanda03dev/go-ms-template/src/core/application/services"
)

type Handlers struct {
	TemplateEntityHandler TemplateEntityHandler
}

var (
	once        sync.Once
	AllHandlers *Handlers
)

func GetHandlers() *Handlers {
	once.Do(func() {
		var AllServices = services.GetServices()
		AllHandlers = &Handlers{
			TemplateEntityHandler: NewTemplateEntityHandler(AllServices.TemplateEntityService),
		}
	})
	return AllHandlers
}
