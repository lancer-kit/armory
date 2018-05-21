package payment

import "gitlab.inn4science.com/vcg/go-common/types/currency"

type (
	AffiliateInfo struct {
		AffiliateID      string `json:"affiliateId"`      // String(50);
		CampaignID       string `json:"campaignId"`       // String(50);
		BannerID         string `json:"bannerId"`         // String(50);
		CustomParameters string `json:"customParameters"` // String(255);
	}

	Parameter struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	Tax struct {
		Name string         `json:"name"` // String(255); Tax name
		Rate currency.Price `json:"rate"` // BigDecimal(3,3); Tax rate. Max value 100
	}

	Transaction struct {
		UID               string        `json:"uid"`
		Type              TxType        `json:"type"`              // String(1); Transaction type
		Action            TxAction      `json:"action"`            // String(2); Transaction action
		Status            TxStatus      `json:"status"`            // String(1); Status
		Amount            currency.Fiat `json:"amount"`            // BigDecimal(14,2); Transaction amount
		PaymentMethodCode string        `json:"paymentMethodCode"` // String(100); Payment method code
		Payer             Account       `json:"payer"`             // Payer object
		Parameters        []Parameter   `json:"parameters"`        // List of name => value pair; Additional parameters
	}

	OrderItem struct {
		Type            OrderIType    `json:"type"`
		Name            string        `json:"name"`            // String(255); Tax name
		Description     string        `json:"description"`     // String(8000); Description
		PriceUnit       currency.Fiat `json:"priceUnit"`       // BigDecimal(14,2);
		Quantity        currency.Fiat `json:"quantity"`        // BigDecimal(14,2);
		DiscountAmount  currency.Fiat `json:"discountAmount"`  // BigDecimal(14,2); Discount in percent (has more priority than discount amount)
		DiscountPercent currency.Fiat `json:"discountPercent"` // BigDecimal(14,2);
		Tax             *Tax          `json:"tax"`
	}

	Order struct {
		UID                          string         `json:"uid"`                          // String(36); ITH platform’s order unique id
		OrderNumber                  string         `json:"orderNumber"`                  // String(20); Visible order number
		Status                       OrderStatus    `json:"status"`                       // String(2)
		Seller                       *Account       `json:"seller"`                       // Order seller object
		Buyer                        *Account       `json:"buyer"`                        // Order buyer object
		CurrencyCode                 string         `json:"currencyCode"`                 // String(3); ISO currency code
		DiscountAmount               currency.Fiat  `json:"discountAmount"`               // BigDecimal(14,2);
		DiscountPercent              currency.Fiat  `json:"discountPercent"`              // BigDecimal(14,2);
		AmountTotal                  currency.Fiat  `json:"amountTotal"`                  // BigDecimal(14,2); Order total amount
		IssueDate                    string         `json:"issueDate"`                    // String(14); Order issue date. Format - yyyyMMddHHmmss
		DueDate                      string         `json:"dueDate"`                      // String(14); Order due date. Format - yyyyMMddHHmmss
		ExternalOrderID              string         `json:"externalOrderId"`              // Order ID in external system
		Reference                    string         `json:"reference"`                    // Reference
		Note                         string         `json:"note"`                         // Note
		Terms                        string         `json:"terms"`                        // Terms
		ProvisionChannel             string         `json:"provisionChannel"`             // String(6); Provision channel
		AffiliateInfo                *AffiliateInfo `json:"affiliateInfo"`                // Affiliate information
		AcceptPaymentsIfOrderExpired bool           `json:"acceptPaymentsIfOrderExpired"` // Accept payments if order expired
		TaxBeforeDiscount            bool           `json:"taxBeforeDiscount"`            // Tax before discount flag
		TaxInclusive                 bool           `json:"taxInclusive"`                 // Tax inclusive flag
		PaymentPageUrl               string         `json:"paymentPageUrl"`               // External payment page URL
		SuccessUrl                   string         `json:"successUrl"`                   // Success URL
		FailUrl                      string         `json:"failUrl"`                      // Fail URL
		OrderItems                   []OrderItem    `json:"orderItems"`                   // Order items
		ShippingAddress              *Address       `json:"shippingAddress"`              // Shipping address
	}

	OrderShort struct {
		ExternalOrderID string        `json:"externalOrderId"` // Merchant’s order id
		UID             string        `json:"uid"`             // String(36); ITH platform’s order unique id
		Seller          *Account      `json:"seller"`          // Order seller object
		Buyer           *Account      `json:"buyer"`           // Order buyer object
		AmountTotal     currency.Fiat `json:"amountTotal"`     // BigDecimal(14,2); Order total amount
		CurrencyCode    string        `json:"currencyCode"`    // String(3); ISO currency code
		Status          OrderStatus   `json:"status"`          // String(2)
		MerchantUrl     string        `json:"merchantUrl"`     // URL for callback sending. Provided only in callbacks
		Transactions    []Transaction `json:"transactions"`    // List of related transactions. Provided only in callbacks
		TestOrder       bool          `json:"testOrder"`
	}
)
