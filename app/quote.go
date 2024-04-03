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
