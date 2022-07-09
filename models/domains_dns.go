package models

import "encoding/xml"

/*
	domains.dns.setDefault
*/
type DomainsDNSSetDefaultResponse struct {
	XMLName           *xml.Name                            `xml:"ApiResponse"`
	ApiResponseStatus *string                              `xml:"Status,attr"`
	Errors            *Error                               `xml:"Errors>Error"`
	CommandResponse   *DomainsDNSSetDefaultCommandResponse `xml:"CommandResponse"`
}

type DomainsDNSSetDefaultCommandResponse struct {
	DomainDNSSetDefaultResult *DomainDNSSetDefaultResultResponse `xml:"DomainDNSSetDefaultResult"`
}

type DomainDNSSetDefaultResultResponse struct {
	Domain  *string `xml:"Domain,attr"`
	Updated *bool   `xml:"Updated,attr"`
}

/*
	domains.dns.setCustom
*/
type DomainsDNSSetCustomResponse struct {
	XMLName           *xml.Name                           `xml:"ApiResponse"`
	ApiResponseStatus *string                             `xml:"Status,attr"`
	Errors            *Error                              `xml:"Errors>Error"`
	CommandResponse   *DomainsDNSSetCustomCommandResponse `xml:"CommandResponse"`
}

type DomainsDNSSetCustomCommandResponse struct {
	DomainDNSSetCustomResult *DomainDNSSetCustomResultResponse `xml:"DomainDNSSetCustomResult"`
}

type DomainDNSSetCustomResultResponse struct {
	Domain  *string `xml:"Domain,attr"`
	Updated *bool   `xml:"Updated,attr"`
}

/*
	domains.dns.getList
*/
type DomainsDNSGetListResponse struct {
	XMLName           *xml.Name                         `xml:"ApiResponse"`
	ApiResponseStatus *string                           `xml:"Status,attr"`
	Errors            *Error                            `xml:"Errors>Error"`
	CommandResponse   *DomainsDNSGetListCommandResponse `xml:"CommandResponse"`
}

type DomainsDNSGetListCommandResponse struct {
	DomainDNSGetListResult *DomainDNSGetListResultResponse `xml:"DomainDNSGetListResult"`
}

type DomainDNSGetListResultResponse struct {
	Domain         *string        `xml:"Domain,attr"`
	IsUsingOurDNS  *bool          `xml:"IsUsingOurDNS,attr"`
	IsPremiumDNS   *bool          `xml:"IsPremiumDNS,attr"`
	IsUsingFreeDNS *bool          `xml:"IsUsingFreeDNS,attr"`
	Nameservers    *[]NameServers `xml:"Nameserver"`
}

type NameServers struct {
	NameServer *string `xml:",chardata"`
}
