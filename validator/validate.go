package validator

import (
	"errors"
	"reflect"
	"strings"
	"unicode"

	"github.com/muathendirangu/convertible/currency"
)

func ValidateInput(c currency.CurrencyRequest) error {
	if  len(strings.TrimSpace(c.From)) == 0 {
		return errors.New("choose the currency you are converting from")
	}else if  len(strings.TrimSpace(c.To)) == 0 {
		return errors.New("choose the currency you are converting to")
	}else if  !isUpper(strings.TrimSpace(c.From)) || !isUpper(strings.TrimSpace(c.To)) {
		return errors.New("currency choice need to be in uppercase format i.e NGN,KES,GHC")
	}else if   c.Amount <= float64(0) {
		return errors.New("amount cannot be zero or a negative number")
	}else if  reflect.TypeOf(c.Amount).Kind() != reflect.Float64  {
		return errors.New("amount must be of type float")
	}
	return nil
}

func isUpper(s string) bool {
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}