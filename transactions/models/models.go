package models

import "time"

/* type TransactionType string

const (
	Income   TransactionType = "income"
	Expense  TransactionType = "expense"
	Transfer TransactionType = "transfer"
) */

type Transaction struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Currensy    string    `json:"currency"`
	Type        string    `json:"type"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

type CreateResponse struct {
	Transaction Transaction `json:"transaction"`
	Status      string      `json:"status"`
}
