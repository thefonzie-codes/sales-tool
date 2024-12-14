package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func main(){
    connectToSqlite()
    insertSQL()
}

func connectToSqlite() {
    db, err := sql.Open("sqlite3", "./sqlite3.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Successfully connected to database.")

    
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS customers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
        account_name TEXT,
        company TEXT,
        email TEXT,
        phone TEXT,
        address TEXT
    );
    `

    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Table 'customer' created or already exists.")
}

func insertSQL () {
    db, err := sql.Open("sqlite3", "./sqlite3.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Successfully connected to database.")

    insertNewCustomerSQL := `
    INSERT INTO customers (
        first_name,
        last_name,
        account_name,
        company,
        email, 
        phone,
        address
    )
    VALUES (?, ?, ?, ?, ?, ?, ?);`

    _, err = db.Exec(
        insertNewCustomerSQL,
        "John",
        "Doe",
        "John Doe",
        "ACME Corp",
        "john@acme.com",
        "(123)-456-7890",
        "1234 Acme Lane, New York, New York",
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Inserted Customer Data")
}
