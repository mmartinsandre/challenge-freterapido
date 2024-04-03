package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func StoreQuoteInDatabase(quote QuoteResponse) error {
	db, err := sql.Open("sqlite3", "./quotes.db")
	if err != nil {
		return err
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
		return err
	}

	for _, carrier := range quote.Carrier {
		_, err = db.Exec("INSERT INTO quotes (carrier_name, service, deadline, price) VALUES (?, ?, ?, ?)",
			carrier.Name, carrier.Service, carrier.Deadline, carrier.Price)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	http.HandleFunc("/quote", QuoteHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
