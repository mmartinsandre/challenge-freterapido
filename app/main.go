package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type QuoteRequest struct {
	Recipient struct {
		Address struct {
			Zipcode string `json:"zipcode"`
		} `json:"address"`
	} `json:"recipient"`
	Volumes []struct {
		Category              int     `json:"category"`
		Amount                int     `json:"amount"`
		UnitaryWeight         int     `json:"unitary_weight"`
		Price                 float64 `json:"price"`
		SKU                   string  `json:"sku"`
		Height, Width, Length float64
	} `json:"volumes"`
}

type QuoteResponse struct {
	Carrier []struct {
		Name     string  `json:"name"`
		Service  string  `json:"service"`
		Deadline string  `json:"deadline"`
		Price    float64 `json:"price"`
	} `json:"carrier"`
}

func QuoteHandler(w http.ResponseWriter, r *http.Request) {
	var reqData QuoteRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

		db, err := sql.Open("sqlite3", "./quotes.db")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS quotes (
								id INTEGER PRIMARY KEY AUTOINCREMENT,
								carrier_name TEXT,
								service TEXT,
								deadline INTEGER,
								price REAL
							)`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, carrier := range quote.Carrier {
			_, err = db.Exec("INSERT INTO quotes (carrier_name, service, deadline, price) VALUES (?, ?, ?, ?)",
				carrier.Name, carrier.Service, carrier.Deadline, carrier.Price)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	quote := QuoteResponse{
		Carrier: []struct {
			Name     string  `json:"name"`
			Service  string  `json:"service"`
			Deadline string  `json:"deadline"`
			Price    float64 `json:"price"`
		}{
			{
				Name:     "EXPRESSO FR",
				Service:  "Rodoviário",
				Deadline: "3",
				Price:    17,
			},
			{
				Name:     "Correios",
				Service:  "SEDEX",
				Deadline: "1",
				Price:    20.99,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

func main() {
	http.HandleFunc("/quote", QuoteHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
