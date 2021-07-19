package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/muathendirangu/convertible/currency"
)

func ValidateInput(c currency.CurrencyRequest) error {
	if  len(strings.TrimSpace(c.From)) == 0 {
		return errors.New("choose the currency you are converting from")
	}else if  len(strings.TrimSpace(c.To)) == 0 {
		return errors.New("choose the currency you are converting to")
	}else if   c.Amount <= float64(0) {
		return errors.New("amount cannot be zero or a negative number")
	}else if  reflect.TypeOf(c.Amount).Kind() != reflect.Float64  {
		return errors.New("amount must be of type float")
	}
	return nil
}