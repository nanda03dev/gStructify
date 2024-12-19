package workers

import (
	"context"
	"log"
	"time"
)

type Worker struct {
	Name    string
	Handler func(ctx context.Context)
}

func InitializeWorkers(ctx context.Context) {
	workers := []Worker{
		{Name: "CRUD Worker", Handler: StartCRUDEventWorker},
	}

	for _, worker := range workers {
		go startWorkerWithRecovery(ctx, worker)
	}
}

func startWorkerWithRecovery(ctx context.Context, worker Worker) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("%s stopped.\n", worker.Name)
			return
		default:
			func() {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("%s panic recovered: %v. Restarting...\n", worker.Name, r)
					}
				}()
				log.Printf("%s started.\n", worker.Name)
				worker.Handler(ctx)
			}()
		}

		// Optional: Add a small delay before restarting
		time.Sleep(2 * time.Second)
	}
}
