package payment

import (
	"testing"
)

func TestProcessPayment(t *testing.T) {

	payment := Payment{
		CardNumber:  "2424242424242424",
		ExpiryMonth: 12,
		ExpiryYear:  2025,
		Amount:      1000,
		Currency:    "GBP",
		CVV:         "123",
	}
	merchantID := "merchant_123"

	result, err := ProcessPayment(payment, merchantID)
	if err != nil {
		t.Fatalf("Failed to process payment: %v", err)
	}
	paymentID := result.ID

	paymentDetails, err := GetPaymentDetails(paymentID)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if paymentDetails.ID != paymentID {
		t.Errorf("Expected payment ID '%v', got '%v'", paymentID, paymentDetails.ID)
	}
}

func TestProcessPayment_InvalidCard(t *testing.T) {
	payment := Payment{
		CardNumber:  "12345678901234565", // Invalid card number
		ExpiryMonth: 12,
		ExpiryYear:  2025,
		Amount:      1000,
		Currency:    "GBP",
		CVV:         "123",
	}
	merchantID := "merchant_123"

	_, err := ProcessPayment(payment, merchantID)

	if err == nil {
		t.Fatal("Expected error for invalid card number, got nil")
	}
}

func TestGetPaymentDetails(t *testing.T) {
	payment := Payment{
		CardNumber:  "4242424242424242",
		ExpiryMonth: 12,
		ExpiryYear:  2025,
		Amount:      1000,
		Currency:    "GBP",
		CVV:         "123",
	}
	merchantID := "merchant_123"

	// Process the payment to store it
	result, err := ProcessPayment(payment, merchantID)
	if err != nil {
		t.Fatalf("Failed to process payment: %v", err)
	}
	paymentID := result.ID

	paymentDetails, err := GetPaymentDetails(paymentID)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if paymentDetails.ID != paymentID {
		t.Errorf("Expected payment ID '%v', got '%v'", paymentID, paymentDetails.ID)
	}
	// Verify that the card number is masked
	expectedMaskedCardNumber := "**** **** **** 4242"
	if paymentDetails.CardNumber != expectedMaskedCardNumber {
		t.Errorf("Expected masked card number '%v', got '%v'", expectedMaskedCardNumber, paymentDetails.CardNumber)
	}

}
