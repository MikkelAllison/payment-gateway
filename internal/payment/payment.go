package payment

import (
	"errors"
	"github.com/MikkelAllison/payment-gateway/internal/bank"
)

type Payment struct {
	ID          string `json:"id,omitempty"`
	CardNumber  string `json:"card_number"`
	ExpiryMonth int    `json:"expiry_month"`
	ExpiryYear  int    `json:"expiry_year"`
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	// Excluded CVV from JSON serialisation to prevent exposing sensitive data.
	CVV string `json:"-"`
}

type PaymentResult struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

var paymentStore = make(map[string]Payment)

// ProcessPayment processes a payment request and communicates with the bank simulator.
func ProcessPayment(p Payment, merchantID string) (PaymentResult, error) {
	// Validate that the card number is 16 digits long.
	if len(p.CardNumber) != 16 {
		return PaymentResult{}, errors.New("invalid card number")
	}

	// Simulate bank processing
	bankResponse, err := bank.ProcessBankPayment(bank.Payment{
		CardNumber:  p.CardNumber,
		ExpiryMonth: p.ExpiryMonth,
		ExpiryYear:  p.ExpiryYear,
		Amount:      p.Amount,
		Currency:    p.Currency,
		CVV:         p.CVV,
	})
	if err != nil {
		return PaymentResult{}, err
	}

	paymentID := bankResponse.TransactionID
	p.ID = paymentID

	paymentStore[paymentID] = p

	result := PaymentResult{
		ID:     paymentID,
		Status: bankResponse.Status,
	}
	return result, nil
}

// GetPaymentDetails retrieves the details of a processed payment by its ID.
func GetPaymentDetails(paymentID string) (Payment, error) {
	payment, exists := paymentStore[paymentID]
	if !exists {
		return Payment{}, errors.New("payment not found")
	}
	// Mask card number
	payment.CardNumber = maskCardNumber(payment.CardNumber)
	return payment, nil
}

// Mask the card number except for the last 4 digits for security reasons.
func maskCardNumber(cardNumber string) string {
	masked := "**** **** **** " + cardNumber[len(cardNumber)-4:]
	return masked
}
