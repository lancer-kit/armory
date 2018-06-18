package uis

import "gitlab.inn4science.com/vcg/go-common/ith/ams"

//Result of request user-integration service @ /ith/auth
//Needs to be signed request (see go-common/auth)
//Example of request: `GET` `http:localhost:2094/v1/uis/ith/auth
//
//   Header. jwt:{"jti":"1"}
//
//Response:
//  {
//		"userId":1,
//		"accessToken":"2bc23zsd3ffer4g993d"
//		"status": "SC"
//	}
type UserAuth struct {
	UserId      int64             `json:"userId"`      //User ID
	AccessToken string            `json:"accessToken"` //Token to process data from/to ITH
	Status      ams.AccountStatus `json:"status"`      //ITH user status (see ams.Account)
}
