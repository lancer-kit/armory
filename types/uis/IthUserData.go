package uis

type IthUserData struct {
	Id        int64  `json:"id"`
	AccountId string `json:"accountId" db:"account_id"`
	AccessToken string `json:"accessToken" db:"access_token"`
	AccountType string `json:"accountType" db:"account_type"`
	AffilateId string `json:"affilateId" db:"affilate_id"`
	CampaignId string `json:"campaignId" db:"campaign_id"`
	BannerId string `json:"bannerId" db:"banner_id"`
	CustomParameters string `json:"customParameters" db:"custom_parameters"`
	AccountSecret string `json:"accountSecret" db:"account_secret"`
	MerchantUid string `json:"merchantUid" db:"merchant_uid"`
	AccountTimezone int `json:"accountTimezone" db:"account_timezone"`
	AccountWeekStartsOn string `json:"accountWeekStartsOn" db:"account_week_starts_on"`
	CurrencyConversion bool `json:"currencyConversion" db:"currency_conversion"`
	AlwaysRefundEwallet bool `json:"alwaysRefundEwallet" db:"always_refund_ewallet"`
	ConfirmOutTransaction bool

}
