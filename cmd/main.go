package main

import (
	"log"

	"github.com/golang_auth/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Database Connection Failed...", err)
	}
}
