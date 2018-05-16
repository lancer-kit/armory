package models

import "gitlab.inn4science.com/vcg/go-common/types/currency"

type (
	Parameter struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	Payer struct {
		Email string `json:"email"`
	}
	Transaction struct {
		UID               string        `json:"uid"`
		Type              string        `json:"type"`
		Status            string        `json:"status"`
		Amount            currency.Fiat `json:"amount"`
		PaymentMethodCode string        `json:"paymentMethodCode"`
		Payer             Payer         `json:"payer"`
		Parameters        []Parameter   `json:"parameters"`
	}

	OrderShort struct {
		ExternalOrderID string        `json:"externalOrderId"`
		UID             string        `json:"uid"`
		AmountTotal     currency.Fiat `json:"amountTotal"`
		CurrencyCode    string        `json:"currencyCode"`
		Status          OrderStatus   `json:"status"` //todo: add status type
		Transactions    []Transaction `json:"transactions"`
		TestOrder       bool          `json:"testOrder"`
		MerchantUrl     string        `json:"merchantUrl"`
	}
)