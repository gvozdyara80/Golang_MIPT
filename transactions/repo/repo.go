package repo

import (
	"errors"
	"sync"
	"transactions/models"
)

var (
	errNotFound = errors.New("transaction not found")
)

type InMemoryDB struct {
	transactions map[string]models.Transaction
	mu           sync.RWMutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		transactions: make(map[string]models.Transaction),
	}
}

func (db *InMemoryDB) Create(transaction models.Transaction) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.transactions[transaction.ID] = transaction
}

func (db *InMemoryDB) Read(id string) (*models.Transaction, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	transaction, exist := db.transactions[id]
	if !exist {
		return nil, errNotFound
	}

	return &transaction, nil
}

func (db *InMemoryDB) Update(id string, updatedTransaction models.Transaction) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, exist := db.transactions[id]
	if !exist {
		return errNotFound
	}

	transaction := updatedTransaction
	db.transactions[id] = transaction

	return nil
}

func (db *InMemoryDB) Delete(id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, exist := db.transactions[id]
	if !exist {
		return errNotFound
	}

	delete(db.transactions, id)

	return nil
}

func (db *InMemoryDB) List() []models.Transaction {
	db.mu.Lock()
	defer db.mu.Unlock()

	transactions := make([]models.Transaction, 0, len(db.transactions))

	for _, transaction := range db.transactions {
		transactions = append(transactions, transaction)
	}
	return transactions
}
