package bank

import (
	"math/rand"
	"time"
)

type Payment struct {
	CardNumber  string
	ExpiryMonth int
	ExpiryYear  int
	Amount      int
	Currency    string
	CVV         string
}

type BankResponse struct {
	TransactionID string
	Status        string
}

func ProcessBankPayment(p Payment) (BankResponse, error) {
	// Simulate processing time
	time.Sleep(100 * time.Millisecond)

	// In this simulation, all bank payments are automatically approved.
	// In a real-world scenario, the bank might decline payments based on various criteria.
	return BankResponse{
		TransactionID: generateTransactionID(),
		Status:        "approved",
	}, nil
}

// Helper function to generate a unique transaction ID
func generateTransactionID() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return "txn_" + randomString(10, r)
}

func randomString(length int, r *rand.Rand) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}
