package db

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlDB struct {
	DB *gorm.DB
}

type MigrationMetadata struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `gorm:"unique;not null"`
	ExecutedAt time.Time `gorm:"autoCreateTime"`
}

func ConnectSqlDB(uri string) (*SqlDB, error) {
	sqlDB := &SqlDB{}
	err := sqlDB.Connect(uri)
	if err != nil {
		return nil, err
	}

	// Run model migrations before starting the application
	if err := RunModelMigration(sqlDB.DB); err != nil {
		log.Fatalf("Could not run model migrations: %v", err)
	}

	// Run raw migrations before starting the application
	if err := RunRawMigration(sqlDB.DB); err != nil {
		log.Fatalf("Could not run raw migrations: %v", err)
	}

	// Run raw migrations before starting the application
	if err := RunSeedMigration(sqlDB.DB); err != nil {
		log.Fatalf("Could not run seed migrations: %v", err)
	}

	// Set connection configs
	sqlDB.setConfigs()

	return sqlDB, nil
}

func (p *SqlDB) Connect(uri string) error {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to SQL Database: %w", err)
	}
	p.DB = db
	return nil
}

func (p *SqlDB) Disconnect() {
	db, err := p.DB.DB()
	if err != nil {
		log.Printf("Error getting raw database connection: %v", err)
		return
	}
	db.Close()
}

func (p *SqlDB) Health() (string, bool) {
	db, err := p.DB.DB()
	if err != nil {
		return "Failed to get raw database connection", false
	}

	db.SetConnMaxLifetime(2 * time.Second)

	err = db.Ping()
	if err != nil {
		return "SQL Database health check failed", false
	} else {
		return "SQL Database connection is healthy!", true
	}
}

func (p *SqlDB) setConfigs() {
	// Get the generic database connection object `*sql.DB` to configure it
	sqlDB, err := p.DB.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	// Set database connection pool parameters
	sqlDB.SetMaxIdleConns(10)               // Maximum number of idle connections
	sqlDB.SetMaxOpenConns(100)              // Maximum number of open connections
	sqlDB.SetConnMaxLifetime(1 * time.Hour) // Maximum time a connection can be reused
}

func RunModelMigration(db *gorm.DB) error {
	// AutoMigrate all registered entities

	if err := db.AutoMigrate(&MigrationMetadata{}); err != nil {
		panic("failed to migrate database")
	}

	for _, entity := range entity.Entities {
		if err := db.AutoMigrate(entity); err != nil {
			panic("failed to migrate database")
		}
	}
	return nil
}

func RunRawMigration(db *gorm.DB) error {
	path := "/sql-migrations/sql/"
	return RunRawSQL(db, path)
}

func RunSeedMigration(db *gorm.DB) error {
	path := "/sql-migrations/seed/"
	return RunRawSQL(db, path)
}

func RecordMigration(db *gorm.DB, migrationName string) error {
	metadata := MigrationMetadata{
		Name: migrationName,
	}
	return db.Create(&metadata).Error
}

func IsMigrationExecuted(db *gorm.DB, migrationName string) (bool, error) {
	var count int64
	err := db.Model(&MigrationMetadata{}).Where("name = ?", migrationName).Count(&count).Error
	return count > 0, err
}

func RunRawSQL(db *gorm.DB, path string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %w", err)
	}

	migrationPath := wd + path

	files, err := os.ReadDir(migrationPath)
	if err != nil {
		// fmt.Println("failed to read migration directory: %w", err)
		return nil
	}

	for _, file := range files {
		if !file.IsDir() {
			migrationFileName := file.Name()

			executed, err := IsMigrationExecuted(db, migrationFileName)
			if err != nil {
				return err
			}

			if executed {
				continue
			}

			// Open the migration file
			f, err := os.Open(migrationPath + migrationFileName)
			if err != nil {
				return fmt.Errorf("failed to open migration file %s: %w", migrationFileName, err)
			}
			defer f.Close()

			// Read the content of the migration file
			script, err := io.ReadAll(f)
			if err != nil {
				return fmt.Errorf("failed to read migration file %s: %w", migrationFileName, err)
			}

			// Execute the migration script
			if err := db.Exec(string(script)).Error; err != nil {
				return fmt.Errorf("failed to execute migration %s: %w", migrationFileName, err)
			}

			// Record the executed migration.
			if err := RecordMigration(db, migrationFileName); err != nil {
				return err
			}
			log.Printf("Executed migration: %s", file.Name())
		}
	}

	return nil
}
