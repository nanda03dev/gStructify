package handlers

import (
	"sync"

	"github.com/nanda03dev/go-ms-template/src/core/application/services"
)

type Handlers struct {
	TemplateEntityHandler TemplateEntityHandler
}

var (
	HandlersOnce sync.Once
	allHandlers  *Handlers
)

func GetHandlers() *Handlers {
	HandlersOnce.Do(func() {
		var AllServices = services.GetServices()
		allHandlers = &Handlers{
			TemplateEntityHandler: NewTemplateEntityHandler(AllServices.TemplateEntityService),
		}
	})
	return allHandlers
}
