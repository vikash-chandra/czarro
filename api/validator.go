package api

import (
	"github.com/czarro/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); !ok {
		// validate currency
		return util.IsSupportCurrency(currency)
	}
	return false
}
