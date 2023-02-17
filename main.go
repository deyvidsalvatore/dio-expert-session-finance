package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main(){

	http.HandleFunc("/", getTransactions)
	http.ListenAndServe(":8000", nil)
}

type Transaction struct {
	Title string
	Amount float32
	Type int
	CreatedAt time.Time
}

type Transactions []Transaction

func getTransactions(w http.ResponseWriter, r *http.Request) {
	
	// Not Allowed Message
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Data Format - YYYY-MM-DDTHH:MM:SS
	var layout = "2006-01-20T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2020-04-05T11:45:26")

	var transactions = Transactions{
		Transaction{
			Title: "Salário",
			Amount: 1200.8,
			Type: 0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}