package handlers

import (
	"encoding/json"
	"net/http"
	"transactions/models"
	"transactions/repo"
)

var InMemoryDB = repo.NewInMemoryDB()

func Transaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var transaction models.Transaction
		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		InMemoryDB.Create(transaction)

		json.NewEncoder(w).Encode(models.TransactionResponse{Transaction: transaction, Status: "succes"})

	case "GET":
		id := r.URL.Query().Get("id")

		if id != "" {
			transaction, err := InMemoryDB.Read(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			json.NewEncoder(w).Encode(models.TransactionResponse{Transaction: *transaction, Status: "succes"})
		} else {
			transactions := InMemoryDB.List()

			json.NewEncoder(w).Encode(models.ListResponse{Transaction: transactions, Status: "succes"})
		}

	case "PUT":
		id := r.URL.Query().Get("id")
		var updTransaction models.Transaction
		if err := json.NewDecoder(r.Body).Decode(&updTransaction); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		InMemoryDB.Update(id, updTransaction)

		json.NewEncoder(w).Encode(models.TransactionResponse{Transaction: updTransaction, Status: "succes"})

	case "DELETE":
		id := r.URL.Query().Get("id")

		InMemoryDB.Delete(id)
		transactions := InMemoryDB.List()

		json.NewEncoder(w).Encode(models.ListResponse{Transaction: transactions, Status: "succes"})
	}
}
