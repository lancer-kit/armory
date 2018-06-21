package uis

import "gitlab.inn4science.com/vcg/go-common/ith/ams"

//ITH -specific user data
//swagger:model
type IthUserData struct {
	Id                        int64                  `json:"id"`                                                //User ID.Same in user-api service
	AccountId                 string                 `json:"accountId" db:"account_id"`                         //External account uid
	AccessToken               string                 `json:"accessToken" db:"access_token"`                     //Refresh token
	AccountType               ams.AccountType        `json:"accountType" db:"account_type"`                     //See ams.AccountType
	AccountStatus             ams.AccountStatus      `json:"accountStatus" db:"account_status"`                 //See ams.AccountStatus
	AffilateId                string                 `json:"affilateId" db:"affiliate_id"`                      //Affilate id (referral)  in ITH AMS
	Language                  string                 `json:"language" db:"language"`                            //String(2). User language ISO2 code
	CommunicationLanguage     string                 `json:"communicationLanguage" db:"communication_language"` //String(2). Communication language
	CampaignId                string                 `json:"campaignId" db:"campaign_id"`                       //Campaign Id (referral)  in ITH AMS
	BannerId                  string                 `json:"bannerId" db:"banner_id"`                           //Banner Id (referral) in ITH AMS
	CustomParameters          string                 `json:"customParameters" db:"custom_parameters"`           //Custom parameters in ITH AMS
	AccountSecret             string                 `json:"accountSecret" db:"account_secret"`                 //Merchant account in ITH AMS
	MerchantUid               string                 `json:"merchantUid" db:"merchant_uid"`                     //Merchant UID in ITH AMS
	AccountTimezone           int                    `json:"accountTimezone" db:"account_timezone"`             //Timezone
	AccountWeekStartsOn       string                 `json:"accountWeekStartsOn" db:"account_week_starts_on"`
	CurrencyConversion        bool                   `json:"currencyConversion" db:"currency_conversion"`
	AlwaysRefundEwallet       bool                   `json:"alwaysRefundEwallet" db:"always_refund_ewallet"`
	ConfirmOutTransaction     bool                   `json:"confirmOutTransaction" db:"confirm_out_transaction"`
	ConfirmLogin              bool                   `json:"confirmLogin" db:"confirm_login"`
	ActionConfirmationEnabled bool                   `json:"actionConfirmationEnabled" db:"action_confirmation_enabled"`
	ActionConfirmationType    ams.ActionConfirmation `json:"actionConfirmationType" db:"action_confirmation_type"`
	Test                      *bool                  `json:"test,omitempty" db:"test"`
}

func GetBoolPtr(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}
