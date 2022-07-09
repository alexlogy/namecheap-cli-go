package common

import (
	"encoding/xml"
	"log"
	"namecheap-cli/consts"
	"namecheap-cli/models"
	"namecheap-cli/utils"
	"strconv"
)

/*
	Only allow Single-domain (DV) for Positive SSL
	and Wildcard (DV) for PositiveSSL Wildcard at the moment
*/
func SSLCreate(certType string, years int) {
	payload := make(map[string]string)
	payload["Command"] = "namecheap.ssl.create"
	payload["Years"] = strconv.Itoa(years)
	payload["Type"] = certType

	var sslCreateResponse models.SSLCreateResponse

	resp := utils.PostRequest(consts.NamecheapAPI, payload)

	err := xml.Unmarshal(resp, &sslCreateResponse)
	if err != nil {
		log.Fatalln(err)
	}

	// api status validation
	respStatus := *sslCreateResponse.ApiResponseStatus

	if respStatus == "ERROR" {
		log.Fatalf("[error] code: %s - message: %s", *sslCreateResponse.Errors.Number, *sslCreateResponse.Errors.Message)
	}

	sslCreateResult := *sslCreateResponse.CommandResponse.SSLCreateResult
	utils.PrintSSLCreateTable(sslCreateResult)
}

/*
	Returns a list of SSL certificates for a particular user
*/
func SSLList() {
	payload := make(map[string]string)
	payload["Command"] = "namecheap.ssl.getList"
	payload["PageSize"] = "100"
	payload["SortBy"] = "EXPIREDATETIME_DESC"

	var sslGetListResponse models.SSLGetListResponse

	resp := utils.PostRequest(consts.NamecheapAPI, payload)

	err := xml.Unmarshal(resp, &sslGetListResponse)
	if err != nil {
		log.Fatalln(err)
	}

	// api status validation
	respStatus := *sslGetListResponse.ApiResponseStatus

	if respStatus == "ERROR" {
		log.Fatalf("[error] code: %s - message: %s", *sslGetListResponse.Errors.Number, *sslGetListResponse.Errors.Message)
	}

	sslGetListResult := *sslGetListResponse.CommandResponse.SSLListResult
	utils.PrintSSLGetListTable(sslGetListResult.SSL)
}

/*
	Activate SSL Cert
*/
func SSLActivate(certID int, csrFile string) {
	payload := make(map[string]string)
	payload["Command"] = "namecheap.ssl.activate"
	payload["CertificateID"] = strconv.Itoa(certID)
	payload["CSR"] = utils.ReadCSRFile(csrFile)
	payload["AdminEmailAddress"] = consts.EmailAddress
	payload["ApproverEmail"] = consts.EmailAddress
	payload["DNSDCValidation"] = "true"

	var sslActivateResponse models.SSLActivateResponse

	resp := utils.PostRequest(consts.NamecheapAPI, payload)

	err := xml.Unmarshal(resp, &sslActivateResponse)
	if err != nil {
		log.Fatalln(err)
	}

	// api status validation
	respStatus := *sslActivateResponse.ApiResponseStatus

	if respStatus == "ERROR" {
		log.Fatalf("[error] code: %s - message: %s", *sslActivateResponse.Errors.Number, *sslActivateResponse.Errors.Message)
	}

	utils.PrintSSLActivateTable(sslActivateResponse)
}
