package db

import (
	"database/sql"
	"log"
	// "os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func ConnectDB() *sql.DB {
    connStr := ""
    // connStr := os.Getenv("DB")
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