package models

import "encoding/xml"

type DomainsGetListResponse struct {
	XMLName           *xml.Name                      `xml:"ApiResponse"`
	ApiResponseStatus *string                        `xml:"Status,attr"`
	Errors            *Error                         `xml:"Errors>Error"`
	CommandResponse   *DomainsGetListCommandResponse `xml:"CommandResponse"`
}

type DomainsGetListCommandResponse struct {
	Domains *[]Domain             `xml:"DomainGetListResult>Domain"`
	Paging  *DomainsGetListPaging `xml:"Paging"`
}

type DomainsGetListPaging struct {
	TotalItems  *int `xml:"TotalItems"`
	CurrentPage *int `xml:"CurrentPage"`
	PageSize    *int `xml:"PageSize"`
}

type Domain struct {
	ID         *string   `xml:"ID,attr"`
	Name       *string   `xml:"Name,attr"`
	User       *string   `xml:"User,attr"`
	Created    *DateTime `xml:"Created,attr"`
	Expires    *DateTime `xml:"Expires,attr"`
	IsExpired  *bool     `xml:"IsExpired,attr"`
	IsLocked   *bool     `xml:"IsLocked,attr"`
	AutoRenew  *bool     `xml:"AutoRenew,attr"`
	WhoisGuard *string   `xml:"WhoisGuard,attr"`
	IsPremium  *bool     `xml:"IsPremium,attr"`
	IsOurDNS   *bool     `xml:"IsOurDNS,attr"`
}

type DomainsCreateResponse struct {
	XMLName           *xml.Name                     `xml:"ApiResponse"`
	ApiResponseStatus *string                       `xml:"Status,attr"`
	Errors            *Error                        `xml:"Errors>Error"`
	CommandResponse   *DomainsCreateCommandResponse `xml:"CommandResponse"`
}

type DomainsCreateCommandResponse struct {
	DomainCreateResult *DomainCreateResultResponse `xml:"DomainCreateResult"`
}

type DomainCreateResultResponse struct {
	TransactionID     *string  `xml:"TransactionID,attr"`
	OrderID           *string  `xml:"OrderID,attr"`
	DomainID          *string  `xml:"DomainID,attr"`
	Domain            *string  `xml:"Domain,attr"`
	Registered        *string  `xml:"Registered,attr"`
	ChargedAmount     *float64 `xml:"ChargedAmount,attr"`
	WhoisguardEnable  *string  `xml:"WhoisguardEnable,attr"`
	NonRealTimeDomain *string  `xml:"NonRealTimeDomain,attr"`
}

/*
domain.setcontacts
*/
type DomainsSetContactsResponse struct {
	XMLName           *xml.Name                          `xml:"ApiResponse"`
	ApiResponseStatus *string                            `xml:"Status,attr"`
	Errors            *Error                             `xml:"Errors>Error"`
	CommandResponse   *DomainsSetContactsCommandResponse `xml:"CommandResponse"`
}

type DomainsSetContactsCommandResponse struct {
	DomainSetContactsResult *DomainSetContactsResultResponse `xml:"DomainSetContactResult"`
}

type DomainSetContactsResultResponse struct {
	Domain    *string `xml:"Domain,attr"`
	IsSuccess *bool   `xml:"IsSuccess,attr"`
}

/*
domain.check
*/
type DomainsCheckResponse struct {
	XMLName           *xml.Name                    `xml:"ApiResponse"`
	ApiResponseStatus *string                      `xml:"Status,attr"`
	Errors            *Error                       `xml:"Errors>Error"`
	CommandResponse   *DomainsCheckCommandResponse `xml:"CommandResponse"`
}

type DomainsCheckCommandResponse struct {
	DomainCheckResult *[]DomainCheckResultResponse `xml:"DomainCheckResult"`
}

type DomainCheckResultResponse struct {
	Domain                   *string  `xml:"Domain,attr"`
	Available                *bool    `xml:"Available,attr"`
	IsPremiumName            *bool    `xml:"IsPremiumName,attr"`
	PremiumRegistrationPrice *float64 `xml:"PremiumRegistrationPrice,attr"`
	PremiumRenewalPrice      *float64 `xml:"PremiumRenewalPrice,attr"`
	PremiumRestorePrice      *float64 `xml:"PremiumRestorePrice,attr"`
	PremiumTransferPrice     *float64 `xml:"PremiumTransferPrice,attr"`
	IcannFee                 *float64 `xml:"IcannFee,attr"`
	EapFee                   *float64 `xml:"EapFee,attr"`
}

/*
domain.renew
*/
type DomainsRenewResponse struct {
	XMLName           *xml.Name                    `xml:"ApiResponse"`
	ApiResponseStatus *string                      `xml:"Status,attr"`
	Errors            *Error                       `xml:"Errors>Error"`
	CommandResponse   *DomainsRenewCommandResponse `xml:"CommandResponse"`
}

type DomainsRenewCommandResponse struct {
	DomainRenewResult *DomainRenewResultResponse `xml:"DomainRenewResult"`
}

type DomainRenewResultResponse struct {
	DomainName    *string        `xml:"DomainName,attr"`
	DomainID      *int           `xml:"DomainID,attr"`
	Renew         *bool          `xml:"Renew,attr"`
	OrderID       *int           `xml:"OrderID,attr"`
	TransactionID *int           `xml:"TransactionID,attr"`
	ChargedAmount *float64       `xml:"ChargedAmount,attr"`
	DomainDetails *DomainDetails `xml:"DomainDetails"`
}

type DomainsReactivateResponse struct {
	XMLName           *xml.Name                         `xml:"ApiResponse"`
	ApiResponseStatus *string                           `xml:"Status,attr"`
	Errors            *Error                            `xml:"Errors>Error"`
	CommandResponse   *DomainsReactivateCommandResponse `xml:"CommandResponse"`
}

type DomainsReactivateCommandResponse struct {
	DomainRenewResult *DomainReactivateResultResponse `xml:"DomainReactivateResult"`
}
type DomainReactivateResultResponse struct {
	DomainName    *string  `xml:"Domain,attr"`
	IsSuccess     *bool    `xml:"IsSuccess,attr"`
	ChargedAmount *float64 `xml:"ChargedAmount,attr"`
	OrderID       *int     `xml:"OrderID,attr"`
	TransactionID *int     `xml:"TransactionID,attr"`
}

type DomainDetails struct {
	ExpiredDate *string `xml:"ExpiredDate"`
	NumYears    *int    `xml:"NumYears"`
}
