package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func fetchCustomer(w http.ResponseWriter, r *http.Request) {

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

func createTable () {
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

func insertCustomer(db *sql.DB, customer Customer) (int64, error) {
	query := `
INSERT INTO customers (first_name, last_name, account_name, company, email, phone, address)
VALUES (?, ?, ?, ?, ?, ?, ?)
`
	result, err := db.Exec(query, customer.FirstName, customer.LastName, customer.AccountName, customer.Company, customer.Email, customer.Phone, customer.Address)
	if err != nil {
		return 0, fmt.Errorf("failed to insert customer: %v", err)
	}

	newID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last inserted ID: %v", err)
	}

	return newID, nil
}

func createCustomerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newCustomer Customer
	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "./sqlite3.db")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect to database: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	newID, err := insertCustomer(db, newCustomer)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert customer: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": newID})
}
