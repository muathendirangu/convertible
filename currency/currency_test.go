package currency

import (
	"fmt"
	"testing"
)

func TestCurrencyExchange(t *testing.T)  {
	rates := RateDictionary{
		"ngnkes":{"NGN","KES",0.26},
		"ngnghc":{"NGN","GHC",0.014},
		"ghckes":{"GHC","KES",18.17},
		"ghcngn":{"GHC","NGN",69.06},
		"kesngn":{"KES","NGN",3.80},
		"kesghc":{"KES","GHC",0.055},
	}
	ex := &ExchangeRate{
		Rates: rates,
	}
	assertMessage := func(t testing.TB, got, want interface{}) {
        t.Helper()
        if got != want {
            t.Errorf("got %f want %f", got, want)
        }
    }
	t.Run("test NGN to NGN",func(t *testing.T) {
		nairaToNaira := ex.Exchange(100,"NGN","NGN")
		want := float64(0)
		got := nairaToNaira.ConvertedAmount
		if got != want {
			fmt.Printf("got %v but want %v", got, want)
		}
	})



	t.Run("test conversion from NGN to KES", func(t *testing.T) {
		nairaToKes :=ex.Exchange(200,"NGN","KES")
		want := float64(52)
		got := nairaToKes.ConvertedAmount
		assertMessage(t,got,want)
	})
	t.Run("test conversion from NGN to GHC", func(t *testing.T) {
		nairaToCedi :=ex.Exchange(200,"NGN","GHC")
		want := float64(2.8)
		got := nairaToCedi.ConvertedAmount
		assertMessage(t,got,want)
	})
	t.Run("test conversion from KES to NGN", func(t *testing.T) {
		kesToNaira :=ex.Exchange(200,"KES","NGN")
		want := float64(760)
		got := kesToNaira.ConvertedAmount
		assertMessage(t,got,want)
	})
	t.Run("test conversion from KES to GHC", func(t *testing.T) {
		kesToCedi :=ex.Exchange(200,"KES","GHC")
		want := float64(11)
		got := kesToCedi.ConvertedAmount
		assertMessage(t,got,want)
	})
	t.Run("test conversion from GHC to NGN", func(t *testing.T) {
		cediToNaira :=ex.Exchange(200,"GHC","NGN")
		want := float64(13812)
		got := cediToNaira.ConvertedAmount
		assertMessage(t,got,want)
	})
	t.Run("test conversion from GHC to KES", func(t *testing.T) {
		cediToKes :=ex.Exchange(200,"GHC","KES")
		want := float64(3634)
		got := cediToKes.ConvertedAmount
		assertMessage(t,got,want)
	})
	t.Run("test passing empty first currency choice", func(t *testing.T) {
		emptyFirstCurrency := ex.Exchange(200,"","NGN")
		want:= float64(0)
		got := emptyFirstCurrency.ConvertedAmount
		assertMessage(t,got,want)
	})
	t.Run("test passing empty second currency choice", func(t *testing.T) {
		emptySecondCurrency := ex.Exchange(200,"NGN","")
		want:= float64(0)
		got := emptySecondCurrency.ConvertedAmount
		assertMessage(t,got,want)
	})
	t.Run("test passing second currency choice in lowercase", func(t *testing.T) {
		emptySecondCurrency := ex.Exchange(200,"NGN","kes")
		want:= "something is off kindly ensure that data input is correct. i.e the conversion currency choice should be uppercase e.g NGN"
		got := emptySecondCurrency.Message
		assertMessage(t,got,want)
	})
	t.Run("test passing amount as zero", func(t *testing.T) {
		emptySecondCurrency := ex.Exchange(0,"NGN","KES")
		want:= float64(0)
		got := emptySecondCurrency.ConvertedAmount
		assertMessage(t,got,want)
	})
}

