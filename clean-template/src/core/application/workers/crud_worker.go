package workers

import (
	"context"
	"log"
	"time"

	"github.com/nanda03dev/go-ms-template/src/core/application/services"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/worker_channels"
)

// StartCRUDWorker listens to the CRUD channel and processes data
func StartCRUDEventWorker(ctx context.Context) {
	for {
		select {
		case event := <-worker_channels.GetCRUDEventChannel(): // Listen to the channel
			log.Println("Processing data:", event)
			// Simulate processing, e.g., saving to database
			if event.Config.EventStore {
				services.NewEventService().Create(event)
			}
			// Optional: Add any delay or logic as required
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			log.Println("Shutting down CRUD Event worker...")
			return
		}
	}
}
