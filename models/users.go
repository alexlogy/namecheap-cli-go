package models

import "encoding/xml"

/*
	Model for users.getBalance API Response in XML
*/
type UsersGetBalanceResponse struct {
	XMLName           *xml.Name                       `xml:"ApiResponse"`
	ApiResponseStatus *string                         `xml:"Status,attr"`
	Errors            *Error                          `xml:"Errors>Error"`
	CommandResponse   *UsersGetBalanceCommandResponse `xml:"CommandResponse"`
}

type UsersGetBalanceCommandResponse struct {
	UserGetBalancesResult *UsersGetBalanceResultResponse `xml:"UserGetBalancesResult"`
}

type UsersGetBalanceResultResponse struct {
	Currency                  *string  `xml:"Currency,attr"`
	AvailableBalance          *float64 `xml:"AvailableBalance,attr"`
	AccountBalance            *float64 `xml:"AccountBalance,attr"`
	EarnedAmount              *float64 `xml:"EarnedAmount,attr"`
	WithdrawableAmount        *float64 `xml:"WithdrawableAmount,attr"`
	FundsRequiredForAutoRenew *float64 `xml:"FundsRequiredForAutoRenew,attr"`
}
