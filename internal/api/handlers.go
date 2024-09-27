package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"

	"github.com/MikkelAllison/payment-gateway/internal/payment"
)

func ProcessPaymentHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var p payment.Payment
	err = json.Unmarshal(body, &p)
	if err != nil {
		http.Error(w, "Invalid payment data", http.StatusBadRequest)
		return
	}

	// Get Merchant-ID from headers (assuming authentication is not implemented)
	merchantID := r.Header.Get("Merchant-ID")
	if merchantID == "" {
		http.Error(w, "Merchant-ID header is required", http.StatusBadRequest)
		return
	}

	result, err := payment.ProcessPayment(p, merchantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)
}

func GetPaymentDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID := vars["id"]

	// Retrieve payment details
	paymentDetails, err := payment.GetPaymentDetails(paymentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Return the payment details as JSON
	json.NewEncoder(w).Encode(paymentDetails)
}
