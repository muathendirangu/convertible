package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/muathendirangu/convertible/currency"
)
	 

func handler(w http.ResponseWriter, r *http.Request) {
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
    req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("an error %v occurred while reading the request body", err)
	}
	json.Unmarshal(req,&currencyReq)
	if err != nil {
		log.Fatalf("an error %v occurred while parsing the request body", err)
	}
	
	conversion := ex.Exchange(currencyReq.Amount,currencyReq.From,currencyReq.To)

	if err != nil {
		log.Fatalf("an error occured %s",err)
	}
	json.NewEncoder(w).Encode(conversion)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}