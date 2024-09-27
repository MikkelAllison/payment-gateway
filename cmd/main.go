package main

import (
	"github.com/MikkelAllison/payment-gateway/internal/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Define API routes
	router.HandleFunc("/payments", api.ProcessPaymentHandler).Methods("POST")
	router.HandleFunc("/payments/{id}", api.GetPaymentDetailsHandler).Methods("GET")

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
