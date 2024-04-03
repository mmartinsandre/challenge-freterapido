package main

import (
	"log"
	"net/http"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./quotes.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT carrier_name, service, deadline, price FROM quotes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var quotes []QuoteResponse
	for rows.Next() {
		var carrierName, service string
		var deadline int
		var price float64
		err := rows.Scan(&carrierName, &service, &deadline, &price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		quotes = append(quotes, QuoteResponse{
			Carrier: []struct {
				Name     string  `json:"name"`
				Service  string  `json:"service"`
				Deadline string  `json:"deadline"`
				Price    float64 `json:"price"`
			}{
				{
					Name:     carrierName,
					Service:  service,
					Deadline: strconv.Itoa(deadline),
					Price:    price,
				},
			},
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}
