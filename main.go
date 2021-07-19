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
		{CurrencyFrom:"NGN",CurrencyTo:"KES",ConversionRate:  0.26},
		{CurrencyFrom:"NGN",CurrencyTo:"GHC",ConversionRate:0.014},
		{CurrencyFrom:"GHC",CurrencyTo:"KES",ConversionRate:18.17},
		{CurrencyFrom:"GHC",CurrencyTo:"NGN",ConversionRate:69.06},
		{CurrencyFrom:"KES",CurrencyTo:"NGN",ConversionRate:3.80},
		{CurrencyFrom:"KES",CurrencyTo:"GHC",ConversionRate:0.055},
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
	log.Printf("Now listening for incoming requests on port %s...\n", os.Getenv("PORT"))
    log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}