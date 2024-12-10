package models

import "time"

// Transaction representa una transacción de débito/crédito
type Transaction struct {
	ID        string    `json:"id"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"`
	Date      time.Time `json:"date"`
	AccountID string    `json:"account_id"`
}
