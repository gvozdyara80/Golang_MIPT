package handlers

import (
	"github.com/Golang_MIPT/8_conversion/currency"
	"github.com/Golang_MIPT/8_conversion/db"
	"github.com/Golang_MIPT/8_conversion/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)


// HandleTransactions роутер для обработки запросов к транзакциям
func HandleTransactions(w http.ResponseWriter, r *http.Request) {
	if currency := r.URL.Query().Get("currency"); currency != "" {
		getTransactionsWithCurrency(w, r)
	} else {
		switch r.Method {
		case "GET":
			getTransactions(w, r)
		case "POST":
			addTransaction(w, r)
		case "PUT":
			updateTransaction(w, r)
		case "DELETE":
			deleteTransaction(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Method not supported"))
		}
	}
}

// addTransaction добавляет новую транзакцию в базу данных
func addTransaction(w http.ResponseWriter, r *http.Request) {
	var t models.Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&t)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

// getTransactions возвращает список всех транзакций из базы данных
func getTransactions(w http.ResponseWriter, _ *http.Request) {
	var transactions []models.Transaction
	result := db.DB.Find(&transactions)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

// updateTransaction обновляет данные транзакции в базе данных
func updateTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var t models.Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t.ID = uint(id)
	result := db.DB.Save(&t)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
}

// deleteTransaction удаляет транзакцию из базы данных
func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	result := db.DB.Delete(&models.Transaction{}, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction deleted"))
}
func CalculateCommission(w http.ResponseWriter, r *http.Request) {
	var c models.Commission
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch {
	case c.Type == "transfer" && c.Currency == "USD":
		c.Commission = c.Amount * 0.02
	case c.Type == "transfer" && c.Currency == "RUB":
		c.Commission = c.Amount * 0.05
	case c.Type == "purchase", c.Type == "top-up":
		c.Commission = 0
	}

	c.Description = "Комиссия за перевод"
	c.Date = time.Now().Format("2024-04-20 15:04:05Z03:00")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func getTransactionsWithCurrency(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transaction
	result := db.DB.Find(&transactions)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	requestedCurrency := r.URL.Query().Get("currency")

	if requestedCurrency != "" {
		currencyClient := currency.NewCurrencyClient("YOUR_API_KEY")
		rates, err := currencyClient.GetCurrencyRate(transactions[0].Currency)
		if err != nil {
			http.Error(w, "Failed to get currency rates", http.StatusInternalServerError)
			return
		}

		convertedTransactions := make([]models.Transaction, 0, len(transactions))
		for _, transaction := range transactions {
			if rate, ok := rates[requestedCurrency]; ok {
				convertedAmount := transaction.Amount * rate
				convertedTransaction := transaction
				convertedTransaction.Converted = &models.ConvertedTransaction{
					Amount:   convertedAmount,
					Currency: requestedCurrency,
				}
				convertedTransactions = append(convertedTransactions, convertedTransaction)
			} else {
				http.Error(w, "Currency not supported", http.StatusBadRequest)
				return
			}
		}

		transactions = convertedTransactions
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
