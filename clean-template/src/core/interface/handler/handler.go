package handler

import (
	"sync"

	"github.com/nanda03dev/go-ms-template/src/core/application/service"
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
		var AllServices = service.GetServices()
		allHandlers = &Handlers{
			TemplateEntityHandler: NewTemplateEntityHandler(AllServices.TemplateEntityService),
		}
	})
	return allHandlers
}
