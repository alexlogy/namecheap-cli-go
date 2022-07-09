package utils

import (
	"fmt"
	tablewriter "github.com/olekukonko/tablewriter"
	"namecheap-cli/models"
	"os"
	"strconv"
	"sync"
)

func WriteTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(data)
	table.SetHeader(header)

	table.Render()
}

func FormatDomainListTable(wg *sync.WaitGroup, output *[][]string, domain models.Domain) {
	defer wg.Done()

	var row []string

	row = []string{
		*domain.ID,
		*domain.Name,
		domain.Created.String(),
		domain.Expires.String(),
		strconv.FormatBool(*domain.IsLocked),
		strconv.FormatBool(*domain.AutoRenew),
		*domain.WhoisGuard,
	}

	*output = append(*output, row)
}

func PrintDomainListTable(domains []models.Domain) {
	var wg sync.WaitGroup
	output := make([][]string, 0, len(domains))
	for _, domain := range domains {
		wg.Add(1)
		go FormatDomainListTable(&wg, &output, domain)
	}

	var header []string
	header = []string{
		"ID",
		"Domain",
		"Created",
		"Expires",
		"IsLocked",
		"AutoRenew",
		"WhoisGuard",
	}

	wg.Wait()

	fmt.Println("Number of domains: ", len(output))
	WriteTable(header, output)
}

func FormatDomainCreateTable(wg *sync.WaitGroup, output *[][]string, domain models.DomainCreateResultResponse) {
	defer wg.Done()

	var row []string
	row = []string{
		*domain.OrderID,
		*domain.TransactionID,
		*domain.DomainID,
		*domain.Domain,
		fmt.Sprintf("%.2f", *domain.ChargedAmount),
		*domain.WhoisguardEnable,
		*domain.NonRealTimeDomain,
	}

	*output = append(*output, row)
}

func PrintDomainCreateTable(domain models.DomainCreateResultResponse) {
	var wg sync.WaitGroup
	output := make([][]string, 0, 1)
	wg.Add(1)
	go FormatDomainCreateTable(&wg, &output, domain)

	var header []string
	header = []string{
		"OrderID",
		"TransactionID",
		"DomainID",
		"Domain",
		"ChargedAmount",
		"Privacy",
		"Real-Time",
	}

	wg.Wait()

	WriteTable(header, output)
}

func FormatDomainCheckTable(wg *sync.WaitGroup, output *[][]string, domain models.DomainCheckResultResponse) {
	defer wg.Done()

	var row []string
	if *domain.IsPremiumName {
		row = []string{
			*domain.Domain,
			strconv.FormatBool(*domain.Available),
			strconv.FormatBool(*domain.IsPremiumName),
			fmt.Sprintf("%.2f", *domain.PremiumRegistrationPrice),
			fmt.Sprintf("%.2f", *domain.PremiumRenewalPrice),
			fmt.Sprintf("%.2f", *domain.PremiumTransferPrice),
		}
	} else {
		// namecheap api doesnt give a price for non-premium name
		row = []string{
			*domain.Domain,
			strconv.FormatBool(*domain.Available),
			strconv.FormatBool(*domain.IsPremiumName),
			"-",
			"-",
			"-",
		}
	}

	*output = append(*output, row)
}

func PrintDomainCheckTable(domains []models.DomainCheckResultResponse) {
	var wg sync.WaitGroup
	output := make([][]string, 0, len(domains))
	for _, domain := range domains {
		wg.Add(1)
		go FormatDomainCheckTable(&wg, &output, domain)
	}

	var header []string
	header = []string{
		"Domain",
		"Available",
		"IsPremiumName",
		"Registration Price",
		"Renewal Price",
		"Transfer Price",
	}

	wg.Wait()

	fmt.Println("Number of domains checked: ", len(output))
	WriteTable(header, output)
}

func FormatDomainDNSSetCustomTable(wg *sync.WaitGroup, output *[][]string, domain models.DomainDNSSetCustomResultResponse, nameServers []string) {
	defer wg.Done()

	var row []string

	row = []string{
		*domain.Domain,
		nameServers[0],
		nameServers[1],
		strconv.FormatBool(*domain.Updated),
	}

	*output = append(*output, row)
}

func PrintDomainDNSSetCustomTable(domain models.DomainDNSSetCustomResultResponse, nameServers []string) {
	var wg sync.WaitGroup
	output := make([][]string, 0, 1)

	wg.Add(1)
	go FormatDomainDNSSetCustomTable(&wg, &output, domain, nameServers)

	var header []string
	header = []string{
		"Domain",
		"NameServer 1",
		"NameServer 2",
		"Updated",
	}

	wg.Wait()

	fmt.Println("Number of domains updated: ", len(output))
	WriteTable(header, output)
}

func FormatDomainDNSSetDefaultTable(wg *sync.WaitGroup, output *[][]string, domain models.DomainDNSSetDefaultResultResponse) {
	defer wg.Done()

	var row []string

	row = []string{
		*domain.Domain,
		"default",
		strconv.FormatBool(*domain.Updated),
	}

	*output = append(*output, row)
}

func PrintDomainDNSSetDefaultTable(domain models.DomainDNSSetDefaultResultResponse) {
	var wg sync.WaitGroup
	output := make([][]string, 0, 1)

	wg.Add(1)
	go FormatDomainDNSSetDefaultTable(&wg, &output, domain)

	var header []string
	header = []string{
		"Domain",
		"NameServers",
		"Updated",
	}

	wg.Wait()

	fmt.Println("Number of domains updated: ", len(output))
	WriteTable(header, output)
}

func FormatDomainDNSGetListTable(wg *sync.WaitGroup, output *[][]string, domain models.DomainDNSGetListResultResponse) {
	defer wg.Done()

	var row []string
	var nameServers []string

	for _, ns := range *domain.Nameservers {
		nameServers = append(nameServers, *ns.NameServer)
	}

	row = []string{
		*domain.Domain,
		strconv.FormatBool(*domain.IsUsingOurDNS),
		strconv.FormatBool(*domain.IsPremiumDNS),
		strconv.FormatBool(*domain.IsUsingFreeDNS),
		nameServers[0],
		nameServers[1],
	}

	*output = append(*output, row)
}

func PrintDomainDNSGetListTable(domain models.DomainDNSGetListResultResponse) {
	var wg sync.WaitGroup
	output := make([][]string, 0, 1)

	wg.Add(1)
	go FormatDomainDNSGetListTable(&wg, &output, domain)

	var header []string
	header = []string{
		"Domain",
		"Namecheap DNS",
		"PremiumDNS",
		"FreeDNS",
		"NameServer 1",
		"NameServer 2",
	}

	wg.Wait()

	fmt.Println("Number of domains checked: ", len(output))
	WriteTable(header, output)
}

func FormatDomainsRenewTable(wg *sync.WaitGroup, output *[][]string,
	domain models.DomainRenewResultResponse, domainDetails models.DomainDetails, years string) {

	defer wg.Done()

	var row []string

	row = []string{
		strconv.Itoa(*domain.OrderID),
		strconv.Itoa(*domain.TransactionID),
		strconv.Itoa(*domain.DomainID),
		*domain.DomainName,
		strconv.FormatBool(*domain.Renew),
		fmt.Sprintf("%.2f", *domain.ChargedAmount),
		years,
		*domainDetails.ExpiredDate,
	}

	*output = append(*output, row)
}

func PrintDomainsRenewTable(domain models.DomainRenewResultResponse, domainDetails models.DomainDetails, years string) {
	var wg sync.WaitGroup
	output := make([][]string, 0, 1)

	wg.Add(1)
	// namecheap API doesn't reflect number of renewal years in API response
	// https://www.namecheap.com/support/api/methods/domains/renew/
	go FormatDomainsRenewTable(&wg, &output, domain, domainDetails, years)

	var header []string
	header = []string{
		"OrderID",
		"TransactionID",
		"DomainID",
		"DomainName",
		"RenewStatus",
		"ChargedAmount",
		"RenewalYrs",
		"New Expiry",
	}

	wg.Wait()

	WriteTable(header, output)
}

func FormatUsersGetBalanceTable(wg *sync.WaitGroup, output *[][]string, users models.UsersGetBalanceResultResponse) {
	defer wg.Done()

	var row []string

	row = []string{
		*users.Currency,
		fmt.Sprintf("%.2f", *users.AvailableBalance),
		fmt.Sprintf("%.2f", *users.AccountBalance),
		fmt.Sprintf("%.2f", *users.EarnedAmount),
		fmt.Sprintf("%.2f", *users.WithdrawableAmount),
		fmt.Sprintf("%.2f", *users.FundsRequiredForAutoRenew),
	}

	*output = append(*output, row)
}

func PrintUsersGetBalanceTable(users models.UsersGetBalanceResultResponse) {
	var wg sync.WaitGroup
	output := make([][]string, 0, 1)

	wg.Add(1)
	go FormatUsersGetBalanceTable(&wg, &output, users)

	var header []string
	header = []string{
		"Currency",
		"AvailableBalance",
		"AccountBalance",
		"EarnedAmount",
		"WithdrawableAmount",
		"FundsRequiredForAutoRenew",
	}

	wg.Wait()

	WriteTable(header, output)
}

func FormatSSLCreateTable(wg *sync.WaitGroup, output *[][]string, ssl models.SSLCreateResultResponse) {
	defer wg.Done()

	var row []string

	row = []string{
		strconv.Itoa(*ssl.OrderId),
		strconv.Itoa(*ssl.TransactionId),
		strconv.Itoa(*ssl.SSLCertificate.CertificateID),
		strconv.FormatBool(*ssl.IsSuccess),
		fmt.Sprintf("%.2f", *ssl.ChargedAmount),
		*ssl.SSLCertificate.Created,
		strconv.Itoa(*ssl.SSLCertificate.Years),
		*ssl.SSLCertificate.Status,
	}

	*output = append(*output, row)
}

func PrintSSLCreateTable(ssl models.SSLCreateResultResponse) {
	var wg sync.WaitGroup
	output := make([][]string, 0, 1)

	wg.Add(1)
	go FormatSSLCreateTable(&wg, &output, ssl)

	var header []string
	header = []string{
		"OrderID",
		"TransactionID",
		"CertID",
		"CreateStatus",
		"ChargedAmount",
		"Created",
		"IssuedYears",
		"Status",
	}

	wg.Wait()

	WriteTable(header, output)
}

func FormatSSLGetListTable(wg *sync.WaitGroup, output *[][]string, ssl models.SSL) {
	defer wg.Done()

	var (
		row              []string
		hostName         string
		expireDate       string
		activationExpire string
	)

	switch *ssl.Status {
	case "newpurchase":
		// newpurchase do not have these fields
		hostName = "-"
		expireDate = "-"
		activationExpire = "-"
	default:
		hostName = *ssl.HostName
		expireDate = *ssl.ExpireDate
		activationExpire = *ssl.ActivationExpireDate
	}

	row = []string{
		strconv.Itoa(*ssl.CertificateID),
		hostName,
		*ssl.SSLType,
		*ssl.PurchaseDate,
		expireDate,
		activationExpire,
		*ssl.Status,
	}

	*output = append(*output, row)
}

func PrintSSLGetListTable(ssl *[]models.SSL) {
	var wg sync.WaitGroup
	output := make([][]string, 0, len(*ssl))

	for _, cert := range *ssl {
		wg.Add(1)
		go FormatSSLGetListTable(&wg, &output, cert)
	}

	var header []string
	header = []string{
		"CertificateID",
		"HostName",
		"SSLType",
		"PurchaseDate",
		"ExpireDate",
		"ActivationExpireDate",
		"Status",
	}

	wg.Wait()

	fmt.Println("Number of SSL Certificates: ", len(output))
	WriteTable(header, output)
}

func FormatSSLActivateTable(wg *sync.WaitGroup, output *[][]string, ssl models.SSLActivateResponse) {
	defer wg.Done()

	var (
		row []string
	)

	row = []string{
		strconv.Itoa(*ssl.CommandResponse.SSLActivateResult.ID),
		*ssl.CommandResponse.SSLActivateResult.DNSDCValidation.DNS.Domain,
		*ssl.CommandResponse.SSLActivateResult.DNSDCValidation.DNS.HostName,
		*ssl.CommandResponse.SSLActivateResult.DNSDCValidation.DNS.Target,
		strconv.FormatBool(*ssl.CommandResponse.SSLActivateResult.DNSDCValidation.ValueAvailable),
	}

	*output = append(*output, row)
}

func PrintSSLActivateTable(ssl models.SSLActivateResponse) {
	var wg sync.WaitGroup
	output := make([][]string, 0, 1)

	wg.Add(1)
	go FormatSSLActivateTable(&wg, &output, ssl)

	var header []string
	header = []string{
		"CertificateID",
		"Domain",
		"DNS Hostname",
		"DNS Target",
		"Activation Status",
	}

	wg.Wait()

	fmt.Println("Number of SSL Certificates Activated: ", len(output))
	WriteTable(header, output)
}
