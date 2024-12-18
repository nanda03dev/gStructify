package db

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func ConnectPostgresDB(uri string) (*PostgresDB, error) {
	postgresDB := &PostgresDB{}
	err := postgresDB.Connect(uri)
	if err != nil {
		return nil, err
	}

	// Run raw migrations before starting the application
	if err := RunRawMigration(postgresDB.DB); err != nil {
		log.Fatalf("Could not run raw migrations: %v", err)
	}

	// Run model migrations before starting the application
	if err := RunModelMigration(postgresDB.DB); err != nil {
		log.Fatalf("Could not run model migrations: %v", err)
	}

	return postgresDB, nil
}

func (p *PostgresDB) Connect(uri string) error {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	p.DB = db
	return nil
}

func (p *PostgresDB) Disconnect() {
	db, err := p.DB.DB()
	if err != nil {
		log.Printf("Error getting raw database connection: %v", err)
		return
	}
	db.Close()
}

func (p *PostgresDB) Health() (string, bool) {
	db, err := p.DB.DB()
	if err != nil {
		return "Failed to get raw database connection", false
	}

	db.SetConnMaxLifetime(2 * time.Second)

	err = db.Ping()
	if err != nil {
		return "PostgreSQL health check failed", false
	} else {
		return "PostgreSQL connection is healthy!", true
	}
}

func RunModelMigration(db *gorm.DB) error {
	// AutoMigrate all registered entities
	for _, entity := range entity.Entities {
		if err := db.AutoMigrate(entity); err != nil {
			panic("failed to migrate database")
		}
	}
	return nil
}

func RunRawMigration(db *gorm.DB) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %w", err)
	}

	migrationPath := wd + "/sql-migrations/"

	files, err := os.ReadDir(migrationPath)
	if err != nil {
		// fmt.Println("failed to read migration directory: %w", err)
		return nil
	}

	for _, file := range files {
		if !file.IsDir() {
			// Open the migration file
			f, err := os.Open(migrationPath + file.Name())
			if err != nil {
				return fmt.Errorf("failed to open migration file %s: %w", file.Name(), err)
			}
			defer f.Close()

			// Read the content of the migration file
			script, err := io.ReadAll(f)
			if err != nil {
				return fmt.Errorf("failed to read migration file %s: %w", file.Name(), err)
			}

			// Execute the migration script
			if err := db.Exec(string(script)).Error; err != nil {
				return fmt.Errorf("failed to execute migration %s: %w", file.Name(), err)
			}
			log.Printf("Executed migration: %s", file.Name())
		}
	}

	return nil
}
