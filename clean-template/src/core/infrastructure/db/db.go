package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type DatabaseInfo struct {
	DB_URI string
}

type Databases struct {
	Postgres *PostgresDB
}

var (
	once               sync.Once
	DatabaseConnection *Databases
	DatabaseURI        DatabaseInfo
)

func init() {
	err := godotenv.Load()

	DatabaseURI = DatabaseInfo{
		DB_URI: "postgres://gtest:gtest@localhost/primary-db?sslmode=disable",
	}

	if err != nil {
		log.Println("No .env file found")
	}

	if DB_URI := os.Getenv("DB_URI"); DB_URI != "" {
		DatabaseURI.DB_URI = DB_URI
	}

}

func ConnectAll() *Databases {
	once.Do(func() {
		postgresDB, err := ConnectPostgresDB(DatabaseURI.DB_URI)
		if err != nil {
			log.Fatalf("Could not connect to PostgreSQL: %v", err)
		} else {
			fmt.Println("successfully connected to PostgreSQL")
		}
		DatabaseConnection = &Databases{
			Postgres: postgresDB,
		}
	})

	return DatabaseConnection
}

func (DatabaseConnection *Databases) DisconnectAll() {
	DatabaseConnection.Postgres.Disconnect()
}
