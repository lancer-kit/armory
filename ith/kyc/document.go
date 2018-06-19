package kyc

import "gitlab.inn4science.com/vcg/go-common/ith"

type (
	Document struct {
		UID            string          `json:"uid,omitempty"`            // String(36); ITH platform’s order unique id
		OwnerUId       string          `json:"ownerUid,omitempty"`       // Owner id
		Status         DocumentStatus  `json:"status,omitempty"`         // String(3)
		Type           DocumentType    `json:"type,omitempty"`           // String(3)
		SubType        DocumentSubType `json:"subtype,omitempty"`        // String(3)
		CountryCode    string          `json:"countryCode,omitempty"`    // String(2)
		Number         string          `json:"number,omitempty"`         // Number of kyc
		UploadDate     ith.Time        `json:"uploadDate,omitempty"`     // yyyyMMDDhhmmss
		ExpirationDate ith.Time        `json:"expirationDate,omitempty"` // yyyyMMDDhhmmss
		FileName1      string          `json:"fileName1,omitempty"`      // string(256)
		DocumentUrl1   string          `json:"documentUrl1,omitempty"`   // string(256); url
		B64File1       string          `json:"base64File1,omitempty"`    // string(256); base64
		FileName2      string          `json:"fileName2,omitempty"`      // string(256)
		DocumentUrl2   string          `json:"documentUrl2,omitempty"`   // string(256); url
		B64File2       string          `json:"base64File2,omitempty"`    // string(256); base64
	}
)
