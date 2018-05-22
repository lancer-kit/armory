package ams

type (
	Account struct {
		Uid                       string             `json:"uid"`                       //Account UID in ITH platform
		Country                   *Country           `json:"country"`                   //Account country object
		Language                  *Language          `json:"language"`                  //Account language object
		CommunicationLanguage     *Language          `json:"communicationLanguage"`     //Account communication language object
		Type                      AccountType        `json:"type"`                      //Account type: see AccountType
		Status                    AccountStatus      `json:"status"`                    //Account status: see AccountStatus
		AccountPhones             AccountPhones      `json:"accountPhones"`             //List of account phones, []*AccountPhone
		AccountSettings           AccountSettings    `json:"accountSettings"`           //List of account settings
		AccountEmails             AccountEmails      `json:"accountEmails"`             //List of account emails
		Addresses                 Addresses          `json:"addresses"`                 //List of Address objects
		Person                    *Person            `json:"person"`                    //Person object
		Company                   *Company           `json:"company,omitempty"`         //Company object (for merchant only)
		AffiliateId               string             `json:"affiliateId"`               //Affiliate ID, String(50)
		CampaignId                string             `json:"campaignId"`                //Campaign ID, String(50)
		BannerId                  string             `json:"bannerId"`                  //Banner ID, String(50)
		CustomParameters          string             `json:"customParameters"`          //Custom parameters, String(255)
		AccountSecret             string             `json:"accountSecret,omitempty"`   //Account secret (for merchant only), String(20)
		MerchantUid               string             `json:"merchantUid,omitempty"`     //Merchant UID (for merchant only), String(36)
		Timezone                  int                `json:"timezone"`                  //Account time zone ID
		WeekStartsOn              string             `json:"weekStartsOn"`              //Start day of the week, String(2)
		CurrencyConversion        bool               `json:"currencyConversion"`        //Currency conversion is enabled
		AlwaysRefundEWallet       bool               `json:"alwaysRefundEWallet"`       //Refunds are transferred to EWallet
		ConfirmOutTransaction     bool               `json:"confirmOutTransaction"`     //2 step verification for outgoing transactions
		ConfirmLogin              bool               `json:"confirmLogin"`              //2 step verification for login
		ActionConfirmationEnabled bool               `json:"actionConfirmationEnabled"` //2 step verification enabled
		ActionConfirmationType    ActionConfirmation `json:"actionConfirmationType"`
		Test                      bool               `json:"test"` //Account is test
	}

	//Standard account response (from doc.example)
	AccountResponse struct {
		Account *Account `json:"account"`
	}
)