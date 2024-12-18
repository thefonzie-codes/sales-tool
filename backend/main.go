package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/add-customer", createCustomerHandler)
	mux.HandleFunc("/customers", fetchCustomer)
	// Enable CORS with default options (allow all origins)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                             // Allows all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},  // Allow these HTTP methods
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Allow these headers
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true, // Allow credentials (cookies, etc.)
	})

	// Wrap your mux with the CORS handler
	handlerWithCORS := c.Handler(mux)

	// Start the server on port 8080
	err := http.ListenAndServe(":8080", handlerWithCORS)
    log.Println("Server started on port 8080")
    if err != nil {
        log.Fatal(http.ListenAndServe(":8080", nil))
    }
}
