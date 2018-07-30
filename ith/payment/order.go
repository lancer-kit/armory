package payment

import (
	"gitlab.inn4science.com/vcg/go-common/ith/auth"
	"gitlab.inn4science.com/vcg/go-common/types/currency"
)

type (
	AffiliateInfo struct {
		AffiliateID      string `json:"affiliateId,omitempty"`      // String(50);
		CampaignID       string `json:"campaignId,omitempty"`       // String(50);
		BannerID         string `json:"bannerId,omitempty"`         // String(50);
		CustomParameters string `json:"customParameters,omitempty"` // String(255);
	}

	Parameter struct {
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	}

	Tax struct {
		Name string         `json:"name,omitempty"` // String(255); Tax name
		Rate currency.Price `json:"rate,omitempty"` // BigDecimal(3,3); Tax rate. Max value 100
	}

	Transaction struct {
		UID               string        `json:"uid,omitempty"`
		Type              TxType        `json:"type,omitempty"`              // String(1); Transaction type
		Action            TxAction      `json:"action,omitempty"`            // String(2); Transaction action
		Status            TxStatus      `json:"status,omitempty"`            // String(1); Status
		Amount            currency.Fiat `json:"amount,omitempty"`            // BigDecimal(14,2); Transaction amount
		PaymentMethodCode string        `json:"paymentMethodCode,omitempty"` // String(100); Payment method code
		Payer             auth.Account  `json:"payer,omitempty"`             // Payer object
		Parameters        []Parameter   `json:"parameters,omitempty"`        // List of name => value pair; Additional parameters
	}

	OrderItem struct {
		Type            OrderIType    `json:"type,omitempty"`
		Name            string        `json:"name,omitempty"`            // String(255); Tax name
		Description     string        `json:"description,omitempty"`     // String(8000); Description
		PriceUnit       currency.Fiat `json:"priceUnit,omitempty"`       // BigDecimal(14,2);
		Quantity        currency.Fiat `json:"quantity,omitempty"`        // BigDecimal(14,2);
		DiscountAmount  currency.Fiat `json:"discountAmount,omitempty"`  // BigDecimal(14,2); Discount in percent (has more priority than discount amount)
		DiscountPercent currency.Fiat `json:"discountPercent,omitempty"` // BigDecimal(14,2);
		Tax             *Tax          `json:"tax,omitempty"`
	}

	Order struct {
		UID                          string         `json:"uid,omitempty"`                          // String(36); ITH platform’s order unique id
		OrderNumber                  string         `json:"orderNumber,omitempty"`                  // String(20); Visible order number
		Status                       OrderStatus    `json:"status,omitempty"`                       // String(2)
		Seller                       *auth.Account  `json:"seller,omitempty"`                       // Order seller object
		Buyer                        *auth.Account  `json:"buyer,omitempty"`                        // Order buyer object
		CurrencyCode                 string         `json:"currencyCode,omitempty"`                 // String(3); ISO currency code
		DiscountAmount               currency.Fiat  `json:"discountAmount,omitempty"`               // BigDecimal(14,2);
		DiscountPercent              currency.Fiat  `json:"discountPercent,omitempty"`              // BigDecimal(14,2);
		AmountTotal                  currency.Fiat  `json:"amountTotal,omitempty"`                  // BigDecimal(14,2); Order total amount
		IssueDate                    string         `json:"issueDate,omitempty"`                    // String(14); Order issue date. Format - yyyyMMddHHmmss
		DueDate                      string         `json:"dueDate,omitempty"`                      // String(14); Order due date. Format - yyyyMMddHHmmss
		ExternalOrderID              string         `json:"externalOrderId,omitempty"`              // Order ID in external system
		Reference                    string         `json:"reference,omitempty"`                    // Reference
		Note                         string         `json:"note,omitempty"`                         // Note
		Terms                        string         `json:"terms,omitempty"`                        // Terms
		ProvisionChannel             string         `json:"provisionChannel,omitempty"`             // String(6); Provision channel
		AffiliateInfo                *AffiliateInfo `json:"affiliateInfo,omitempty"`                // Affiliate information
		AcceptPaymentsIfOrderExpired bool           `json:"acceptPaymentsIfOrderExpired,omitempty"` // Accept payments if order expired
		TaxBeforeDiscount            bool           `json:"taxBeforeDiscount,omitempty"`            // Tax before discount flag
		TaxInclusive                 bool           `json:"taxInclusive,omitempty"`                 // Tax inclusive flag
		PaymentPageUrl               string         `json:"paymentPageUrl,omitempty"`               // External payment page URL
		SuccessUrl                   string         `json:"successUrl,omitempty"`                   // Success URL
		FailUrl                      string         `json:"failUrl,omitempty"`                      // Fail URL
		OrderItems                   []OrderItem    `json:"orderItems,omitempty,omitempty"`         // Order items
		ShippingAddress              *auth.Address  `json:"shippingAddress,omitempty,omitempty"`    // Shipping address
	}

	OrderShort struct {
		ExternalOrderID string        `json:"externalOrderId,omitempty"` // Merchant’s order id
		UID             string        `json:"uid,omitempty"`             // String(36); ITH platform’s order unique id
		Seller          *auth.Account `json:"seller,omitempty"`          // Order seller object
		Buyer           *auth.Account `json:"buyer,omitempty"`           // Order buyer object
		AmountTotal     currency.Fiat `json:"amountTotal,omitempty"`     // BigDecimal(14,2); Order total amount
		CurrencyCode    string        `json:"currencyCode,omitempty"`    // String(3); ISO currency code
		Status          OrderStatus   `json:"status,omitempty"`          // String(2)
		MerchantUrl     string        `json:"merchantUrl,omitempty"`     // URL for callback sending. Provided only in callbacks
		Transactions    []Transaction `json:"transactions,omitempty"`    // List of related transactions. Provided only in callbacks
		TestOrder       bool          `json:"testOrder,omitempty"`
	}
)
