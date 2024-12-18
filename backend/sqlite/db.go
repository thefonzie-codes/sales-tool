package sqlite_connection

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3" // SQLite driver
)

var DB *sql.DB

func Init() {
    var err error
    DB, err = sql.Open("sqlite3", "./cadence.db")
    if err != nil {
        log.Fatal(err)
    }

    // Create tables if they don't exist
    createTableQueries := []string{
        `CREATE TABLE IF NOT EXISTS users (
            id TEXT PRIMARY KEY,
            name TEXT,
            email TEXT,
            password TEXT,
            role TEXT,
            created_at TIMESTAMP,
            updated_at TIMESTAMP
        )`,
        // Add create table queries for other models here
    }

    for _, query := range createTableQueries {
        _, err := DB.Exec(query)
        if err != nil {
            log.Fatal(err)
        }
    }
}

func Close() {
    DB.Close()
}

func GetDB() *sql.DB {
    return DB
}

