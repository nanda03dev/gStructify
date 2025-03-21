package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type DatabaseInfo struct {
	SQL_DB_URI   string
	MONGO_DB_URI string
}

type Databases struct {
	SqlDB   *SqlDB
	MongoDB *MongoDB
}

var (
	once               sync.Once
	DatabaseConnection *Databases
	DatabaseURI        DatabaseInfo
)

func init() {
	err := godotenv.Load()

	DatabaseURI = DatabaseInfo{
		MONGO_DB_URI: "mongodb://localhost:27017",
		SQL_DB_URI:   "postgres://gtest:gtest@host.docker.internal:5432/primary-db?sslmode=disable",
	}

	if err != nil {
		log.Println("No .env file found")
	}

	if SQL_DB_URI := os.Getenv("SQL_DB_URI"); SQL_DB_URI != "" {
		DatabaseURI.SQL_DB_URI = SQL_DB_URI
	}

	if MONGO_DB_URI := os.Getenv("MONGO_DB_URI"); MONGO_DB_URI != "" {
		DatabaseURI.MONGO_DB_URI = MONGO_DB_URI
	}
}

func ConnectAll() *Databases {
	once.Do(func() {
		sqlDB, err := ConnectSqlDB(DatabaseURI.SQL_DB_URI)
		if err != nil {
			log.Fatalf("Could not connect to SQL Database: %v", err)
		} else {
			fmt.Println("successfully connected to SQL Database")
		}

		mongoDB, err := ConnectMongoDB(DatabaseURI.MONGO_DB_URI)

		if err != nil {
			log.Fatalf("Could not connect to MongoDB: %v", err)
		} else {
			fmt.Println("successfully connected to MongoDB")
		}

		DatabaseConnection = &Databases{
			SqlDB:   sqlDB,
			MongoDB: mongoDB,
		}
	})

	return DatabaseConnection
}

func (DatabaseConnection *Databases) DisconnectAll() {
	DatabaseConnection.SqlDB.Disconnect()
}
