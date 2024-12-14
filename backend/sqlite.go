package main

import (
    "database/sql"
    "fmt"
    "log"
    "encoding/json"
    "net/http"

    _ "github.com/mattn/go-sqlite3"
)

type Customer struct {
    ID          int    `json:"id"`
    FirstName   string `json:"first_name"`
    LastName    string `json:"last_name"`
    AccountName string `json:"account_name"`
    Company     string `json:"company"`
    Email       string `json:"email"`
    Phone       string `json:"phone"`
    Address     string `json:"address"`
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

func fetchCustomer (w http.ResponseWriter, r *http.Request) {

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

    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")

    name := r.URL.Query().Get("name")
    if name == "" {
        http.Error(w, "Missing 'name' in request", http.StatusBadRequest)
        return
    }

    var customer Customer
    query := `SELECT * FROM customers WHERE LOWER(first_name) = LOWER(?)`
    err = db.QueryRow(query, name).Scan(
        &customer.ID, &customer.FirstName, &customer.LastName, &customer.AccountName,
        &customer.Company, &customer.Phone, &customer.Email, &customer.Address,
    )	
    if err == sql.ErrNoRows {
        http.Error(w, "Customer not found", http.StatusNotFound)
        return
    } else if err != nil {
        log.Printf("Failed to fetch customer: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    err = json.NewEncoder(w).Encode(customer)
    if err != nil {
        log.Printf("Failed to encode customer data to JSON: %v", err)
        http.Error(w, "Failed to encode customer data", http.StatusInternalServerError)
    }

    // Optionally, you can print the customer to the log if needed for debugging
    fmt.Printf("Fetched customer: %+v\n", customer)
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
