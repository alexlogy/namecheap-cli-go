package models

import "encoding/xml"

/*
	Model for ssl.create API Response in XML
*/
type SSLCreateResponse struct {
	XMLName           *xml.Name                 `xml:"ApiResponse"`
	ApiResponseStatus *string                   `xml:"Status,attr"`
	Errors            *Error                    `xml:"Errors>Error"`
	CommandResponse   *SSLCreateCommandResponse `xml:"CommandResponse"`
}

type SSLCreateCommandResponse struct {
	SSLCreateResult *SSLCreateResultResponse `xml:"SSLCreateResult"`
}

type SSLCreateResultResponse struct {
	IsSuccess      *bool           `xml:"IsSuccess,attr"`
	OrderId        *int            `xml:"OrderId,attr"`
	TransactionId  *int            `xml:"TransactionId,attr"`
	ChargedAmount  *float64        `xml:"ChargedAmount,attr"`
	SSLCertificate *SSLCertificate `xml:"SSLCertificate"`
}

type SSLCertificate struct {
	CertificateID *int    `xml:"CertificateID,attr"`
	Created       *string `xml:"Created,attr"`
	SSLType       *string `xml:"SSLType,attr"`
	Years         *int    `xml:"Years,attr"`
	Status        *string `xml:"Status,attr"`
}

/*
	Model for ssl.getList API Response in XML
*/
type SSLGetListResponse struct {
	XMLName           *xml.Name                  `xml:"ApiResponse"`
	ApiResponseStatus *string                    `xml:"Status,attr"`
	Errors            *Error                     `xml:"Errors>Error"`
	CommandResponse   *SSLGetListCommandResponse `xml:"CommandResponse"`
}

type SSLGetListCommandResponse struct {
	SSLListResult *SSLGetListResultResponse `xml:"SSLListResult"`
}

type SSLGetListResultResponse struct {
	SSL *[]SSL `xml:"SSL"`
}

type SSL struct {
	CertificateID        *int    `xml:"CertificateID,attr"`
	HostName             *string `xml:"HostName,attr"`
	SSLType              *string `xml:"SSLType,attr"`
	PurchaseDate         *string `xml:"PurchaseDate,attr"`
	ExpireDate           *string `xml:"ExpireDate,attr"`
	ActivationExpireDate *string `xml:"ActivationExpireDate,attr"`
	IsExpiredYN          *bool   `xml:"IsExpiredYN,attr"`
	Status               *string `xml:"Status,attr"`
}

/*
	Model for ssl.activate API Response in XML
*/
type SSLActivateResponse struct {
	XMLName           *xml.Name                   `xml:"ApiResponse"`
	ApiResponseStatus *string                     `xml:"Status,attr"`
	Errors            *Error                      `xml:"Errors>Error"`
	CommandResponse   *SSLActivateCommandResponse `xml:"CommandResponse"`
}

type SSLActivateCommandResponse struct {
	SSLActivateResult *SSLActivateResultResponse `xml:"SSLActivateResult"`
}

type SSLActivateResultResponse struct {
	ID              *int                        `xml:"ID,attr"`
	IsSuccess       *bool                       `xml:"IsSuccess,attr"`
	DNSDCValidation *SSLDNSDCValidationResponse `xml:"DNSDCValidation"`
}

type SSLDNSDCValidationResponse struct {
	ValueAvailable *bool `xml:"ValueAvailable,attr"`
	DNS            *DNS  `xml:"DNS"`
}

type DNS struct {
	Domain   *string `xml:"domain,attr"`
	HostName *string `xml:"HostName"`
	Target   *string `xml:"Target"`
}
