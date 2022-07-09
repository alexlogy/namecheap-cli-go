package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"
	"namecheap-cli/consts"
)

/*
	Generate Private Key
*/
func GeneratePrivateKey() *rsa.PrivateKey {
	// generate private key
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		log.Fatalln(err)
	}

	return privatekey
}

/*
	Generate CSR
*/
func GenerateCSR(domainNames []string, emailAddresses []string, privKey *rsa.PrivateKey) []byte {
	subject := pkix.Name{
		Country:            []string{consts.Country},
		Organization:       []string{consts.Organization},
		OrganizationalUnit: []string{consts.OrganizationUnit},
		Locality:           []string{consts.City},
		Province:           []string{consts.StateProvince},
		StreetAddress:      []string{consts.Address1},
		PostalCode:         []string{consts.PostalCode},
		CommonName:         domainNames[0],
	}
	certRequest := x509.CertificateRequest{
		Subject: subject,
	}
	// generate csr
	csr, err := x509.CreateCertificateRequest(rand.Reader, &certRequest, privKey)
	if err != nil {
		log.Fatalln(err)
	}

	return csr
}
