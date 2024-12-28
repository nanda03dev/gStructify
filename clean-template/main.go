package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/nanda03dev/go-ms-template/src/bootstrap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fiberApp := fiber.New()

	applicationManager := bootstrap.NewApplicationManager(ctx, fiberApp)

	applicationManager.ConnectDatabase()
	applicationManager.Run()

	// Handle graceful shutdown
	go func() {
		// Listen for termination signals (Ctrl+C, kill)
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan

		// Graceful shutdown logic
		log.Println("Shutting down worker...")
		cancel() // Stop workers

		// Gracefully shut down the Fiber app
		log.Println("Shutting down Fiber app...")
		fiberApp.Shutdown()
		applicationManager.DisconnectDatabase()
	}()

	// Start listening for HTTP requests
	log.Println("Starting server on :3000")
	if err := fiberApp.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
