package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func DbConnect() (*sql.DB, error) {
	// get env variables
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbNetProtocol := os.Getenv("DB_NET_PROTOCOL")

	// Data source name properties
	dsn := mysql.Config{
		User:   dbUsername,
		Passwd: dbPassword,
		Net:    dbNetProtocol,
		Addr:   dbHost,
		DBName: dbName,
	}

	// Get a database handle
	db, err := sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("[db] connection successful")
	return db, nil
}
