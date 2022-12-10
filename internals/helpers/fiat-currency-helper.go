package fiat_currency

import (
	"hailo/internals/dtos/requests"
	"hailo/internals/models"
)

func FilterAvailableCurrencies(
	coingeckoCurrencies []requests.Currency,
	systemCurrencies map[string]*models.FiatCurrency) []requests.Currency {

	var result []requests.Currency

	for _, item := range coingeckoCurrencies {
		if _, ok := systemCurrencies[item.Code]; ok {
			result = append(result, item)
		}
	}

	return result
}
