package userapi

import "regexp"

type (
	//Request body for PUT v1/api/user-state/
	PhoneSearch struct {
		Phone string `required:"yes" description:"search value: user phone or wallet id" json:"phone" example:"+380555555555"`
	}
	//Response from USER API user-state controller
	UserState struct {
		Id         int    `json:"id" description:"User internal ID" example:"123"`
		WalletId   string `json:"walletId" description:"User wallet ID" example:"grslaktlzhhxp7ye6r2vxl5zvemv6epdbhg6qwgtu7z5orpjtola"`
		Status     int64  `json:"status"  description:"User status code" example:"30"`
		StatusName string `json:"statusName" description:"User status name" example:"user authorized"`
	}

	//No user found data
	ErrResponse struct {
		Errcode int         `json:"errcode"`
		Errmsg  interface{} `json:"errmsg"`
	}
)

var rule = regexp.MustCompile("(^+)|([^a-zA-Z0-9]+)")

func (t *PhoneSearch) CleanPhone() {
	t.Phone = rule.ReplaceAllString(t.Phone, "")
}
