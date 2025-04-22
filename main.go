package main

import (
	"log"

	db "social-network/backend/pkg/db/migrations"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
