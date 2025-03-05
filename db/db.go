package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func ConnectDB() *sql.DB {
    connStr := "host=198.19.249.38 port=5433 user=postgres password=21wqsaXZ dbname=postgres sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Database connected!")
    return db
}