package worker

import (
	"context"
	"log"
	"time"

	"github.com/nanda03dev/go-ms-template/src/core/application/service"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/worker_channel"
)

// StartCRUDWorker listens to the CRUD channel and processes data
func StartCRUDEventWorker(ctx context.Context) {
	crudEventChannel := worker_channel.GetCRUDEventChannel()
	for {
		select {
		case event := <-crudEventChannel: // Listen to the channel
			// Simulate processing, e.g., saving to database
			if event.Config.EventStore {
				AllServices := service.GetServices()
				AllServices.EventService.Create(event)
			}
			// Optional: Add any delay or logic as required
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			log.Println("Shutting down CRUD Event worker...")
			return
		}
	}
}
