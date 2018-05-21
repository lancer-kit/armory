package ams

type (
	StandardAccount struct {
		Uid                   string            `json:"uid"`                   //Account UID in ITH platform
		Country               *Country          `json:"country"`               //Account country object
		Language              *Language         `json:"language"`              //Account language object
		CommunicationLanguage *Language         `json:"communicationLanguage"` //Account communication language object
		Type                  AccountType       `json:"type"`                  //Account type: see AccountType
		Status                AccountStatus     `json:"status"`                //Account status: see AccountStatus
		AccountPhones         []*AccountPhone   `json:"accountPhones"`         //List of account phones
		AccountSettings       []*AccountSetting `json:"accountSettings"`       //List of account settings
		AccountEmails         []*AccountEmail   `json:"accountEmails"`         //List of account emails
		Addresses             []*Address        `json:"addresses"`             //List of Address objects
		Person                *Person           `json:"person"`
		AffiliateId           string            `json:"affiliateId"`
		CampaignId            string            `json:"campaignId"`
		BannerId              string            `json:"bannerId"`
		CustomParameters      string            `json:"customParameters"`
		CurrencyConversion    bool              `json:"currencyConversion"`
		AlwaysRefundEWallet   bool              `json:"alwaysRefundEWallet"`
		ConfirmOutTransaction bool              `json:"confirmOutTransaction"`
		Test                  bool              `json:"test"`
	}
)
