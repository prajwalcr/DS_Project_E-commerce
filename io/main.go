package io

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {

	err := godotenv.Load("io/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	// create database called twopc
	connStr := fmt.Sprintf("user=%s dbname=twopc password=%s host=localhost sslmode=disable", user, password)
	db, err := sql.Open("postgres", connStr)
	DB = db
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	//defer DB.Close()
}
