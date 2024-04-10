package models

import (
	"time"
)

type Commission struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Type          string  `json:"type"`
	Commission    float64 `json:"commission"`
	Date          string  `json:"date"`
	Description   string  `json:"description"`
}

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

type Transaction struct {
	ID          uint                  `gorm:"primaryKey" json:"id"`
	UserID      uint                  `json:"user_id"`
	Amount      float64               `json:"amount"`
	Currency    string                `json:"currency"`
	Type        string                `json:"type"`
	Category    string                `json:"category"`
	Date        time.Time             `json:"date"`
	Description string                `json:"description"`
	Converted   *ConvertedTransaction `json:"converted,omitempty"`
}

type ConvertedTransaction struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
