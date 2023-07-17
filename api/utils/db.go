package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func DbConnect() (db *sql.DB) {
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
	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
		fmt.Println("error is here ")
	}

	defer db.Close()

	pingErr := db.Ping()
	if err != nil {
		log.Fatal(pingErr)
		fmt.Println("error is here at ping ")
	}

	// upon successful connection
	db.SetConnMaxLifetime(10 * time.Second)
	log.Println("[db] connection successful")
	return
}
