package kyc

import (
	"gitlab.inn4science.com/vcg/go-common/ith"
)

const (
	APIDocumentList   = "/commonapi/document/list"
	APIDocumentUpload = "/commonapi/document/upload"
)

type ErrorData struct {
	ith.ErrorData
}

type ListResp struct {
	ErrorData    *ErrorData `json:"errorData,omitempty"`
	DocumentList []Document `json:"documentList,omitempty"`
}

type CreateDocumentRequest struct {
	ErrorData   *ErrorData `json:"errorData,omitempty"`
	Document    *Document  `json:"document,omitempty"`
	RedirectUrl string     `json:"redirectUrl,omitempty"`
}
