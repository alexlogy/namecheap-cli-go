package common

import (
	"encoding/xml"
	"log"
	"namecheap-cli/consts"
	"namecheap-cli/models"
	"namecheap-cli/utils"
	"regexp"
	"strconv"
	"strings"
)

func DomainsGetList(domain string, page int) {
	payload := make(map[string]string)
	payload["Command"] = "namecheap.domains.getList"
	payload["Page"] = strconv.Itoa(page)
	payload["PageSize"] = "100"
	payload["SortBy"] = "EXPIREDATE_DESC"

	if domain != "" {
		// check whether domain is specified
		payload["SearchTerm"] = domain
	}

	var domainGetListResponse models.DomainsGetListResponse

	resp := utils.PostRequest(consts.NamecheapAPI, payload)

	err := xml.Unmarshal(resp, &domainGetListResponse)
	if err != nil {
		log.Fatalln(err)
	}

	// error handling for API response
	respStatus := *domainGetListResponse.ApiResponseStatus

	if respStatus == "ERROR" {
		log.Fatalf("[error] code: %s - message: %s", *domainGetListResponse.Errors.Number, *domainGetListResponse.Errors.Message)
	}

	domains := *domainGetListResponse.CommandResponse.Domains
	totalItems := *domainGetListResponse.CommandResponse.Paging.TotalItems
	//currentPage := *domainGetListResponse.CommandResponse.Paging.CurrentPage

	if totalItems != 0 {
		// print domain table
		utils.PrintDomainListTable(domains)
	} else {
		log.Printf("There are %d domains in the account.", totalItems)
	}
}

func DomainsCreate(domain string) {
	payload := make(map[string]string)
	payload["Command"] = "namecheap.domains.create"
	payload["Years"] = "1" // fix to 1 year
	payload["AddFreeWhoisguard"] = "yes"
	payload["WGEnabled"] = "yes"
	if domain != "" {
		// check for valid domain
		regex := "(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]"
		matched, err := regexp.MatchString(regex, domain)
		if !matched {
			log.Fatalln("[error] invalid domain input!")
		}
		if err != nil {
			log.Fatalln(err)
		}
		payload["DomainName"] = domain
	}

	InitRegistrant() // initialize default registrant information

	payload["RegistrantFirstName"] = registrant.RegistrantFirstName
	payload["RegistrantLastName"] = registrant.RegistrantLastName
	payload["RegistrantAddress1"] = registrant.RegistrantAddress1
	payload["RegistrantCity"] = registrant.RegistrantCity
	payload["RegistrantStateProvince"] = registrant.RegistrantStateProvince
	payload["RegistrantPostalCode"] = registrant.RegistrantPostalCode
	payload["RegistrantCountry"] = registrant.RegistrantCountry
	payload["RegistrantPhone"] = registrant.RegistrantPhone
	payload["RegistrantEmailAddress"] = registrant.RegistrantEmailAddress
	payload["TechFirstName"] = registrant.TechFirstName
	payload["TechLastName"] = registrant.TechLastName
	payload["TechAddress1"] = registrant.TechAddress1
	payload["TechCity"] = registrant.TechCity
	payload["TechStateProvince"] = registrant.TechStateProvince
	payload["TechPostalCode"] = registrant.TechPostalCode
	payload["TechCountry"] = registrant.TechCountry
	payload["TechPhone"] = registrant.TechPhone
	payload["TechEmailAddress"] = registrant.TechEmailAddress
	payload["AdminFirstName"] = registrant.AdminFirstName
	payload["AdminLastName"] = registrant.AdminLastName
	payload["AdminAddress1"] = registrant.AdminAddress1
	payload["AdminCity"] = registrant.AdminCity
	payload["AdminStateProvince"] = registrant.AdminStateProvince
	payload["AdminPostalCode"] = registrant.AdminPostalCode
	payload["AdminCountry"] = registrant.AdminCountry
	payload["AdminPhone"] = registrant.AdminPhone
	payload["AdminEmailAddress"] = registrant.AdminEmailAddress
	payload["AuxBillingFirstName"] = registrant.AuxBillingFirstName
	payload["AuxBillingLastName"] = registrant.AuxBillingLastName
	payload["AuxBillingAddress1"] = registrant.AuxBillingAddress1
	payload["AuxBillingCity"] = registrant.AuxBillingCity
	payload["AuxBillingStateProvince"] = registrant.AuxBillingStateProvince
	payload["AuxBillingPostalCode"] = registrant.AuxBillingPostalCode
	payload["AuxBillingCountry"] = registrant.AuxBillingCountry
	payload["AuxBillingPhone"] = registrant.AuxBillingPhone
	payload["AuxBillingEmailAddress"] = registrant.AuxBillingEmailAddress

	resp := utils.PostRequest(consts.NamecheapAPI, payload)

	var domainCreateResponse models.DomainsCreateResponse
	err := xml.Unmarshal(resp, &domainCreateResponse)
	if err != nil {
		log.Fatalln(err)
	}

	// error handling for API response
	respStatus := *domainCreateResponse.ApiResponseStatus

	if respStatus == "ERROR" {
		log.Fatalf("[error] code: %s - message: %s", *domainCreateResponse.Errors.Number, *domainCreateResponse.Errors.Message)
	}

	isRegistered := *domainCreateResponse.CommandResponse.DomainCreateResult.Registered
	domainCreateResult := *domainCreateResponse.CommandResponse.DomainCreateResult

	if isRegistered != "false" {
		// print domain table
		utils.PrintDomainCreateTable(domainCreateResult)
		log.Println("Domain Successfully Registered.")
	}
}

func DomainSetContact(domain string) {
	payload := make(map[string]string)
	payload["Command"] = "namecheap.domains.setContacts"
	if domain != "" {
		// check for valid domain
		regex := "(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]"
		matched, err := regexp.MatchString(regex, domain)
		if !matched {
			log.Fatalln("[error] invalid domain input!")
		}
		if err != nil {
			log.Fatalln(err)
		}
		payload["DomainName"] = domain
	}

	InitRegistrant() // initialize default registrant information

	payload["RegistrantFirstName"] = registrant.RegistrantFirstName
	payload["RegistrantLastName"] = registrant.RegistrantLastName
	payload["RegistrantAddress1"] = registrant.RegistrantAddress1
	payload["RegistrantCity"] = registrant.RegistrantCity
	payload["RegistrantStateProvince"] = registrant.RegistrantStateProvince
	payload["RegistrantPostalCode"] = registrant.RegistrantPostalCode
	payload["RegistrantCountry"] = registrant.RegistrantCountry
	payload["RegistrantPhone"] = registrant.RegistrantPhone
	payload["RegistrantEmailAddress"] = registrant.RegistrantEmailAddress
	payload["TechFirstName"] = registrant.TechFirstName
	payload["TechLastName"] = registrant.TechLastName
	payload["TechAddress1"] = registrant.TechAddress1
	payload["TechCity"] = registrant.TechCity
	payload["TechStateProvince"] = registrant.TechStateProvince
	payload["TechPostalCode"] = registrant.TechPostalCode
	payload["TechCountry"] = registrant.TechCountry
	payload["TechPhone"] = registrant.TechPhone
	payload["TechEmailAddress"] = registrant.TechEmailAddress
	payload["AdminFirstName"] = registrant.AdminFirstName
	payload["AdminLastName"] = registrant.AdminLastName
	payload["AdminAddress1"] = registrant.AdminAddress1
	payload["AdminCity"] = registrant.AdminCity
	payload["AdminStateProvince"] = registrant.AdminStateProvince
	payload["AdminPostalCode"] = registrant.AdminPostalCode
	payload["AdminCountry"] = registrant.AdminCountry
	payload["AdminPhone"] = registrant.AdminPhone
	payload["AdminEmailAddress"] = registrant.AdminEmailAddress
	payload["AuxBillingFirstName"] = registrant.AuxBillingFirstName
	payload["AuxBillingLastName"] = registrant.AuxBillingLastName
	payload["AuxBillingAddress1"] = registrant.AuxBillingAddress1
	payload["AuxBillingCity"] = registrant.AuxBillingCity
	payload["AuxBillingStateProvince"] = registrant.AuxBillingStateProvince
	payload["AuxBillingPostalCode"] = registrant.AuxBillingPostalCode
	payload["AuxBillingCountry"] = registrant.AuxBillingCountry
	payload["AuxBillingPhone"] = registrant.AuxBillingPhone
	payload["AuxBillingEmailAddress"] = registrant.AuxBillingEmailAddress

	resp := utils.PostRequest(consts.NamecheapAPI, payload)

	var domainSetContactsResponse models.DomainsSetContactsResponse
	err := xml.Unmarshal(resp, &domainSetContactsResponse)
	if err != nil {
		log.Fatalln(err)
	}

	// error handling for API response
	respStatus := *domainSetContactsResponse.ApiResponseStatus

	if respStatus == "ERROR" {
		log.Fatalf("[error] code: %s - message: %s", *domainSetContactsResponse.Errors.Number, *domainSetContactsResponse.Errors.Message)
	}

	isSuccess := *domainSetContactsResponse.CommandResponse.DomainSetContactsResult.IsSuccess

	if isSuccess {
		log.Printf("Contacts for domain (%s) successfully set.", domain)
	}
}

func DomainsCheck(domains []string) {
	// check for invalid domain
	for _, value := range domains {
		// check for valid domain
		regex := "(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]"
		matched, err := regexp.MatchString(regex, value)
		if !matched {
			log.Fatalln("[error] invalid domain input!")
		}
		if err != nil {
			log.Fatalln(err)
		}
	}
	payload := make(map[string]string)
	payload["Command"] = "namecheap.domains.check"
	payload["DomainList"] = strings.Join(domains, ",")

	var domainCheckResponse models.DomainsCheckResponse

	resp := utils.PostRequest(consts.NamecheapAPI, payload)

	err := xml.Unmarshal(resp, &domainCheckResponse)
	if err != nil {
		log.Fatalln(err)
	}

	// error handling for api response
	respStatus := *domainCheckResponse.ApiResponseStatus

	if respStatus == "ERROR" {
		log.Fatalf("[error] code: %s - message: %s", *domainCheckResponse.Errors.Number, *domainCheckResponse.Errors.Message)
	}

	domainsCheckResults := *domainCheckResponse.CommandResponse.DomainCheckResult
	utils.PrintDomainCheckTable(domainsCheckResults)
}

func DomainsRenew(domain string, years string) {
	// check for invalid domain
	regex := "(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]"
	matched, err := regexp.MatchString(regex, domain)
	if !matched {
		log.Fatalln("[error] invalid domain input!")
	}
	if err != nil {
		log.Fatalln(err)
	}

	payload := make(map[string]string)
	payload["Command"] = "namecheap.domains.renew"
	payload["DomainName"] = domain
	payload["Years"] = years

	var domainRenewResponse models.DomainsRenewResponse

	resp := utils.PostRequest(consts.NamecheapAPI, payload)

	err = xml.Unmarshal(resp, &domainRenewResponse)
	if err != nil {
		log.Fatalln(err)
	}

	// api status validation
	respStatus := *domainRenewResponse.ApiResponseStatus

	if respStatus == "ERROR" {
		log.Fatalf("[error] code: %s - message: %s", *domainRenewResponse.Errors.Number, *domainRenewResponse.Errors.Message)
	}

	domainsRenewResults := *domainRenewResponse.CommandResponse.DomainRenewResult
	domainDetails := *domainRenewResponse.CommandResponse.DomainRenewResult.DomainDetails

	utils.PrintDomainsRenewTable(domainsRenewResults, domainDetails, years)
}
