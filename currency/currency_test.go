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
	ex := &ExchangeRate{
		Rates: rates,
	}
	assertMessage := func(t testing.TB, got, want interface{}) {
        t.Helper()
        if got != want {
            t.Errorf("got %f want %f", got, want)
        }
    }
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
	// t.Run("test passing negative amount", func(t *testing.T) {
	// 	passNegativeAmount := ex.Exchange(-200,"NGN","KES")
	// 	want:= "amount can only be of type float and must be a number greater than zero"
	// 	got := passNegativeAmount.Message
	// 	assertMessage(t,got,want)
	// })
	// t.Run("test passing both currency choices with empty values", func(t *testing.T) {
	// 	emptyBothCurrency := ex.Exchange(200,"","")
	// 	want:= float64(0)
	// 	got := emptyBothCurrency.ConvertedAmount
	// 	assertMessage(t,got,want)
	// })
	// t.Run("test passing zero values", func(t *testing.T) {
	// 	passZeroValues := ex.Exchange(0,"","")
	// 	want:= "currency choices or amount cannot be blank"
	// 	got := passZeroValues.Message
	// 	assertMessage(t,got,want)
	// })
}

