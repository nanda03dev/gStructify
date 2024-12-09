package db

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type DatabaseInfo struct {
	DB_URI string
}

type Databases struct {
	DB *Database
}

var (
	once               sync.Once
	DatabaseConnection *Databases
	DatabaseURI        DatabaseInfo
)

func init() {
	err := godotenv.Load()

	DatabaseURI = DatabaseInfo{
		DB_URI: "mongodb://localhost:27017",
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
		// Declare Database connection
		// mongoDB, err := ConnectMongoDB(DatabaseURI.DB_URI)
		// DatabaseConnection = &Databases{
		// 	DB: mongoDB,
		// }
	})

	return DatabaseConnection
}

func (DatabaseConnection *Databases) DisconnectAll() {
	DatabaseConnection.DB.Disconnect()
}

type Database struct {
	DB_URI string
	// Client *mongo.Database
}

func NewDatabase(DB_URI string) *Database {
	return &Database{DB_URI: DB_URI}
}

func (d *Database) Connect(uri string) error {
	// Implement Database connection logic
	// m.Client = client.Database("go-ms-template")

	return nil
}

func (d *Database) Disconnect() {
	// Implement Database connection logic
	// d.Disconnect(context.Background())
}

func (d *Database) Health() (string, bool) {
	// Implement Health connection logic
	return "Connected to database successfully!", true
}
