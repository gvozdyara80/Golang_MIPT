package models

import "time"

type Transaction struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Currensy    string    `json:"currency"`
	Type        string    `json:"type"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

type TransactionResponse struct {
	Transaction Transaction `json:"transaction"`
	Status      string      `json:"status"`
}

type ListResponse struct {
	Transaction []Transaction `json:"transaction"`
	Status      string        `json:"status"`
}
