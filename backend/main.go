package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/customers", fetchCustomer)

	log.Println("Server us running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
