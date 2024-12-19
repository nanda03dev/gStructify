package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/nanda03dev/go-ms-template/src/core/application/workers"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/core/interface/router"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	var databases = db.ConnectAll()

	defer databases.DisconnectAll()

	app := fiber.New()

	router.InitializeRoutes(app)

	// Initialize application modules
	log.Println("Starting workers...")
	workers.InitializeWorkers(ctx)

	// Handle graceful shutdown
	go func() {
		// Listen for termination signals (Ctrl+C, kill)
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan

		// Graceful shutdown logic
		log.Println("Shutting down workers...")
		cancel() // Stop workers

		// Gracefully shut down the Fiber app
		log.Println("Shutting down Fiber app...")
		app.Shutdown()
	}()

	// Start listening for HTTP requests
	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
