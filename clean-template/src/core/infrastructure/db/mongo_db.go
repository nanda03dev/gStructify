package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	DB *mongo.Database
}

func ConnectMongoDB(uri string) (*MongoDB, error) {
	mongoDB := &MongoDB{}
	err := mongoDB.Connect(uri)
	if err != nil {
		return nil, err
	}
	return mongoDB, nil
}

func (m *MongoDB) Connect(uri string) error {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatalf("Error while initializing MongoDB client: %v", err)
	}

	// Ping to check connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}

	m.DB = client.Database("go-ms-template")

	return nil
}

func (m *MongoDB) Disconnect() {
	m.DB.Client().Disconnect(context.Background())
}

func (m *MongoDB) Health() (string, bool) {
	// Define a context with a timeout for the ping operation
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Run a ping to the MongoDB server
	err := m.DB.Client().Ping(ctx, nil)

	if err != nil {
		return "MongoDB health check failed", false
	} else {
		return "MongoDB connection is healthy!", true
	}

}
