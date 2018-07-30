package kyc

//go:generate goplater -type=DocumentType -transform=none -tprefix=false
type DocumentType int

const (
	DocumentTypeIDN DocumentType = 1 + iota
	DocumentTypeUBL
	DocumentTypeSLF
)

//go:generate goplater -type=DocumentSubType -transform=none -tprefix=false
type DocumentSubType int

const (
	DocumentSubTypeIDL DocumentSubType = 1 + iota
	DocumentSubTypeIRC
	DocumentSubTypeIDC
	DocumentSubTypeIRP
	DocumentSubTypeIPS
	DocumentSubTypeUBL
	DocumentSubTypeUCS
	DocumentSubTypeUHR
	DocumentSubTypeUPL
	DocumentSubTypeUGI
	DocumentSubTypeSLF
)

//go:generate goplater -type=DocumentStatus -transform=none -tprefix=false
type DocumentStatus int

const (
	DocumentStatusNEW DocumentStatus = 1 + iota
	DocumentStatusVLD
	DocumentStatusWAR
	DocumentStatusERR
)
