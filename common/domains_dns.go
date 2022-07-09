package common

import (
	"encoding/xml"
	"log"
	"namecheap-cli/consts"
	"namecheap-cli/models"
	"namecheap-cli/utils"
	"regexp"
	"strings"
)

func DomainsDNSSetDefault(domain string) {
	// check for invalid domain
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
	}

	domainSlice := strings.Split(domain, ".")

	payload := make(map[string]string)
	payload["Command"] = "namecheap.domains.dns.setDefault"
	payload["SLD"] = domainSlice[0]
	payload["TLD"] = domainSlice[1]

	var domainDNSSetDefaultResponse models.DomainsDNSSetDefaultResponse

	resp := utils.PostRequest(consts.NamecheapAPI, payload)

	err := xml.Unmarshal(resp, &domainDNSSetDefaultResponse)
	if err != nil {
		log.Fatalln(err)
	}

	domainsSetDefaultResults := *domainDNSSetDefaultResponse.CommandResponse.DomainDNSSetDefaultResult
	utils.PrintDomainDNSSetDefaultTable(domainsSetDefaultResults)
}

func DomainsDNSSetCustom(domain string, nameServers []string) {
	// check for invalid domain
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
	}

	// check for nameservers
	for _, value := range nameServers {
		regex := "(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]"
		matched, err := regexp.MatchString(regex, value)
		if !matched {
			log.Fatalln("[error] invalid nameserver!")
		}
		if err != nil {
			log.Fatalln(err)
		}
	}

	domainSlice := strings.Split(domain, ".")

	payload := make(map[string]string)
	payload["Command"] = "namecheap.domains.dns.setCustom"
	payload["SLD"] = domainSlice[0]
	payload["TLD"] = domainSlice[1]
	payload["Nameservers"] = strings.Join(nameServers, ",")

	var domainDNSSetCustomResponse models.DomainsDNSSetCustomResponse

	resp := utils.PostRequest(consts.NamecheapAPI, payload)

	err := xml.Unmarshal(resp, &domainDNSSetCustomResponse)
	if err != nil {
		log.Fatalln(err)
	}

	domainsSetCustomResults := *domainDNSSetCustomResponse.CommandResponse.DomainDNSSetCustomResult
	utils.PrintDomainDNSSetCustomTable(domainsSetCustomResults, nameServers)
}

func DomainsDNSGetList(domain string) {
	// check for invalid domain
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
	}

	domainSlice := strings.Split(domain, ".")

	payload := make(map[string]string)
	payload["Command"] = "namecheap.domains.dns.getList"
	payload["SLD"] = domainSlice[0]
	payload["TLD"] = domainSlice[1]

	var domainDNSGetListResponse models.DomainsDNSGetListResponse

	resp := utils.PostRequest(consts.NamecheapAPI, payload)

	err := xml.Unmarshal(resp, &domainDNSGetListResponse)
	if err != nil {
		log.Fatalln(err)
	}

	domainDNSGetListResult := *domainDNSGetListResponse.CommandResponse.DomainDNSGetListResult
	utils.PrintDomainDNSGetListTable(domainDNSGetListResult)
}
