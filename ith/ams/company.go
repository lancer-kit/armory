package ams

//Company
// Company object (for merchant only)
type Company struct {
	Id     int64  `json:"id,omitempty" db:"id"`          //Internal for user-integration
	UserId int64  `json:"userId,omitempty" db:"user_id"` //Internal for user-integration
	Uid    string `json:"uid,omitempty" db:"uid"`        //Internal for user-integration Company UID

	//ITH.AMS data structure
	BusinessName                    string  `json:"businessName" db:"business_name"`                                         //Company name, String(255)
	CategoryId                      int     `json:"categoryId" db:"category_id"`                                             //Category ID, Integer
	BusinessTypeId                  int     `json:"businessTypeId" db:"business_type_id"`                                    //Business type ID, Integer
	CardStatementName               string  `json:"cardStatementName" db:"card_statement_name"`                              //Card statement name, String(50)
	CardStatementNameExt            string  `json:"cardStatementNameExt" db:"card_statement_name_ext"`                       //Extended card statement name, String(50)
	CallbackUrl                     string  `json:"callbackUrl" db:"callback_url"`                                           //URL for callbacks
	RollingReservePrc               float64 `json:"rollingReservePrc" db:"rolling_reserve_prc"`                              //Rolling reserve rate (in %), Number
	RollingReserveHoldDays          int     `json:"rollingReserveHoldDays" db:"rolling_reserve_hold_days"`                   //Rolling reserve hold days
	SendCallback                    bool    `json:"sendCallback" db:"send_callback"`                                         //Send callbacks for merchant
	AcceptUndefinedProvisionChannel bool    `json:"acceptUndefinedProvisionChannel" db:"accept_undefined_provision_channel"` //Accept undefined provision channels
	AllowDuplicateOrderExternalId   bool    `json:"allowDuplicateOrderExternalId" db:"allow_duplicate_order_external_id"`    //Allow duplicate order external ID
	AllowNotificationsForSeller     bool    `json:"allowNotificationsForSeller" db:"allow_notifications_for_seller"`         //Send notifications for seller
	AllowNotificationsForBuyer      bool    `json:"allowNotificationsForBuyer" db:"allow_notifications_for_buyer"`           //Send notifications for buyer
	AllowPartialPayments            bool    `json:"allowPartialPayments" db:"allow_partial_payments"`                        //Allow partial payments
}
