package uis

import "gitlab.inn4science.com/vcg/go-common/ith/ams"

type IthUserData struct {
	Id                        int64                  `json:"id"`
	AccountId                 string                 `json:"accountId" db:"account_id"`
	AccessToken               string                 `json:"accessToken" db:"access_token"`
	AccountType               ams.AccountType        `json:"accountType" db:"account_type"`
	AccountStatus             ams.AccountStatus      `json:"accountStatus" db:"account_status"`
	AffilateId                string                 `json:"affilateId" db:"affilate_id"`
	CampaignId                string                 `json:"campaignId" db:"campaign_id"`
	BannerId                  string                 `json:"bannerId" db:"banner_id"`
	CustomParameters          string                 `json:"customParameters" db:"custom_parameters"`
	AccountSecret             string                 `json:"accountSecret" db:"account_secret"`
	MerchantUid               string                 `json:"merchantUid" db:"merchant_uid"`
	AccountTimezone           int                    `json:"accountTimezone" db:"account_timezone"`
	AccountWeekStartsOn       string                 `json:"accountWeekStartsOn" db:"account_week_starts_on"`
	CurrencyConversion        bool                   `json:"currencyConversion" db:"currency_conversion"`
	AlwaysRefundEwallet       bool                   `json:"alwaysRefundEwallet" db:"always_refund_ewallet"`
	ConfirmOutTransaction     bool                   `json:"confirmOutTransaction" db:"confirm_out_transaction"`
	ConfirmLogin              bool                   `json:"confirmLogin" db:"confirm_login"`
	ActionConfirmationEnabled bool                   `json:"actionConfirmationEnabled" db:"action_confirmation_enabled"`
	ActionConfirmationType    ams.ActionConfirmation `json:"actionConfirmationType" db:"action_confirmation_type"`
	Test                      bool                   `json:"test" db:"test"`
}
