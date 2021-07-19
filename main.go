package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/muathendirangu/convertible/currency"
	"github.com/muathendirangu/convertible/validator"
)
	 

func currencyConverterHandler(w http.ResponseWriter, r *http.Request) {
	rates := []currency.Rate{
		{"NGN","KES",0.26},
		{"NGN","GHC",0.014},
		{"GHC","KES",18.17},
		{"GHC","NGN",69.06},
		{"KES","NGN",3.80},
		{"KES","GHC",0.055},
	}
	ex := currency.ExchangeRate{ 
		Rates: rates,
	}

	currencyReq:= currency.CurrencyRequest{}
	if err := json.NewDecoder(r.Body).Decode(&currencyReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err:=  validator.ValidateInput(currencyReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conversion := ex.Exchange(currencyReq.Amount,currencyReq.From,currencyReq.To)
	json.NewEncoder(w).Encode(conversion)
}

func main() {
	 err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }
    http.HandleFunc("/", currencyConverterHandler)
    log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}