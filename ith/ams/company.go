package ams
//Company
// Company object (for merchant only)
type Company struct {
	BusinessName                    string  `json:"businessName"`                    //Company name, String(255)
	CategoryId                      int     `json:"categoryId"`                      //Category ID, Integer
	BusinessTypeId                  int     `json:"businessTypeId"`                  //Business type ID, Integer
	CardStatementName               string  `json:"cardStatementName"`               //Card statement name, String(50)
	CardStatementNameExt            string  `json:"cardStatementNameExt"`            //Extended card statement name, String(50)
	CallbackUrl                     string  `json:"callbackUrl"`                     //URL for callbacks
	RollingReservePrc               float64 `json:"rollingReservePrc"`               //Rolling reserve rate (in %), Number
	RollingReserveHoldDays          int     `json:"rollingReserveHoldDays"`          //Rolling reserve hold days
	SendCallback                    bool    `json:"sendCallback"`                    //Send callbacks for merchant
	AcceptUndefinedProvisionChannel bool    `json:"acceptUndefinedProvisionChannel"` //Accept undefined provision channels
	AllowDuplicateOrderExternalId   bool    `json:"allowDuplicateOrderExternalId"`   //Allow duplicate order external ID
	AllowNotificationsForSeller     bool    `json:"allowNotificationsForSeller"`     //Send notifications for seller
	AllowNotificationsForBuyer      bool    `json:"allowNotificationsForBuyer"`      //Send notifications for buyer
	AllowPartialPayments            bool    `json:"allowPartialPayments"`            //Allow partial payments
}
