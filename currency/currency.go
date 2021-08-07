package currency

import (
	"fmt"

	"math"
)
type CurrencyExchange interface{
	Exchange(float64,string,string)(CurrencyResponse,error)
}
type CurrencyRequest struct{
	From string `json:"from"`
	To string `json:"to"`
	Amount float64 `json:"amount"`
}

type CurrencyResponse struct{
	From,To string
	InitialAmount,ConvertedAmount float64
	DefaultConversionRate float64
	Message interface{}
	
}
type RateDictionary map[string]Rate

type Rate struct{
	CurrencyFrom string `json:"currencyfrom"`
	CurrencyTo string `json:"currencyto"`
	ConversionRate float64  `json:"conversionrate"`
}

type ExchangeRate struct{
	Rates RateDictionary
}
type Rates interface {
    CurrentConversion() Rate
    SetCurrentConversion(updatedConversion Rate)
}

type CurrentConversion struct {
    currentConversion Rate
}

func (cc CurrentConversion) CurrentConversion() Rate {
    return cc.currentConversion
}

func (cc *CurrentConversion) SetCurrentConversion(updatedConversion Rate) {
    cc.currentConversion = updatedConversion
}

func (e *ExchangeRate) Exchange(amount float64, from,to string)(*CurrencyResponse){
	for _,v := range e.Rates{
		if v.CurrencyFrom == from && v.CurrencyTo==to {
			return &CurrencyResponse{
				From: from,
				To: to,
				InitialAmount: amount,
				ConvertedAmount:math.Round((amount*v.ConversionRate)*100)/100,
				DefaultConversionRate:v.ConversionRate,
				Message: fmt.Sprintf(" Amount %v %s is equivalent to %v %s after conversion using the rate of 1 %v equals %v %v",amount,from,math.Round((amount*v.ConversionRate)*100)/100,to,v.CurrencyFrom,v.ConversionRate,v.CurrencyTo),
				
			} 
		}
	}
	return &CurrencyResponse{From: from,
				To: to,
				InitialAmount: amount,
		Message: "something is off kindly ensure that data input is correct. i.e the conversion currency choice should be uppercase e.g NGN",
	}  
}