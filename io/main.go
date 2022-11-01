package io

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// Change user and password and also create database called twopc
	connStr := "user=vishishtrao dbname=twopc password=12345 host=localhost sslmode=disable"
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
