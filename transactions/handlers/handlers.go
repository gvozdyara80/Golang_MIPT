package handlers

import (
	"encoding/json"
	"net/http"
	"transactions/models"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(models.CreateResponse{Transaction: transaction, Status: "succes"})
}
