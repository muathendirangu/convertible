package currency

import (
	"testing"
)

func TestCurrencyExchange(t *testing.T)  {
	rates := []Rate{
		{"NGN","KES",0.26},
		{"NGN","GHC",0.014},
		{"GHC","KES",18.17},
		{"GHC","NGN",69.06},
		{"KES","NGN",3.80},
		{"KES","GHC",0.055},
	}
	ex := ExchangeRate{
		Rates: rates,
	}

	nairaToKes :=ex.Exchange(200,"NGN","KES")
	
	if nairaToKes.Amount !=52 {
		t.Errorf("Expected 52, but got %f",nairaToKes)
	}
	nairaToCedi :=ex.Exchange(200,"NGN","GHC")
	
	if nairaToCedi.Amount !=2.8 {
		t.Errorf("Expected 2.8, but got %f",nairaToCedi)
	}

	kesToNaira :=ex.Exchange(200,"KES","NGN")
	
	if kesToNaira.Amount !=760 {
		t.Errorf("Expected 760, but got %f",kesToNaira)
	}
	

}