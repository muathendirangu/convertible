package currency

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)
type CurrencyRequest struct{
	From string `json:"from"`
	To string `json:"to"`
	Amount float64 `json:"amount"`
}

type CurrencyResponse struct{
	Message interface{} 
}
type Rate struct{
	CurrencyFrom string `json:"currencyfrom"`
	CurrencyTo string `json:"currencyto"`
	ConversionRate float64  `json:"conversionrate"`
}
type ErrorResponse struct{
	Error error
}

type ExchangeRate struct{
	Rates []Rate
}

func (e ExchangeRate) Exchange(amount float64, from,to string)(interface{}){
	if len(strings.TrimSpace(from))  == 0 || len(strings.TrimSpace(to))  == 0  {
		return &CurrencyResponse{
			Message: " currency choices cannot be blank",
		}
	}else if reflect.TypeOf(amount).Kind()  != reflect.Float64{
		return &CurrencyResponse{
			Message: "amount can only be of type float",
		}
	}
	for _,v := range e.Rates{
		if v.CurrencyFrom == from && v.CurrencyTo==to {
			return &CurrencyResponse{
				Message: fmt.Sprintf(" Amount %v %s is equivalent to %v %s after conversion",amount,from,math.Round((amount*v.ConversionRate)*100)/100,to),
			} 
		}
	}
	return &CurrencyResponse{
		Message: "something is off kindly ensure that data input is correct",
	}  
}