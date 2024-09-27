// internal/bank/bank_test.go
package bank

import (
	"testing"
)

func TestProcessBankPayment_Approved(t *testing.T) {
	payment := Payment{
		CardNumber:  "4242424242424242",
		ExpiryMonth: 12,
		ExpiryYear:  2025,
		Amount:      1000, // Even amount to simulate approval
		Currency:    "GBP",
		CVV:         "123",
	}

	response, err := ProcessBankPayment(payment)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response.Status != "approved" {
		t.Errorf("Expected status 'approved', got '%v'", response.Status)
	}
	if response.TransactionID == "" {
		t.Error("Expected non-empty transaction ID")
	}
}
