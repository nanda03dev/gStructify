package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type DatabaseInfo struct {
	SQL_DB_URI string
}

type Databases struct {
	SqlDB *SqlDB
}

var (
	once               sync.Once
	DatabaseConnection *Databases
	DatabaseURI        DatabaseInfo
)

func init() {
	err := godotenv.Load()

	DatabaseURI = DatabaseInfo{
		SQL_DB_URI: "postgres://gtest:gtest@localhost/primary-db?sslmode=disable",
	}

	if err != nil {
		log.Println("No .env file found")
	}

	if SQL_DB_URI := os.Getenv("SQL_DB_URI"); SQL_DB_URI != "" {
		DatabaseURI.SQL_DB_URI = SQL_DB_URI
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
		DatabaseConnection = &Databases{
			SqlDB: sqlDB,
		}
	})

	return DatabaseConnection
}

func (DatabaseConnection *Databases) DisconnectAll() {
	DatabaseConnection.SqlDB.Disconnect()
}
