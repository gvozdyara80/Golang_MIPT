package cmd

import (
	"log"
	"net/http"
	"transactions/handlers"
)

func Run() {
	http.HandleFunc("/transactions", handlers.Transaction)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
