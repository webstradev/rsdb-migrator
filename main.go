package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/webstradev/rsdb-migrator/importer"
	"github.com/webstradev/rsdb-migrator/migrations"
)

func main() {
	// If a database connection string is not yet set in environment variables (or by kube secrets) then load the .env file
	if os.Getenv("DB_CONNECTION_STRING") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading environment file: %v", err)
		}
	}

	// Import the necessary json files from the mongodumps folder
	data, err := importer.ImportFiles()
	if err != nil {
		log.Fatalf("Error importing	files: %v", err)
	}

	// Create database connection
	db, err := sqlx.Open("mysql", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("Could not open database connection: %v", err)
	}

	// Run Migrations
	migrations.Migrate(db, data)
}
