package commands

import (
	"github.com/spf13/cobra"
	"log"
	"namecheap-cli/consts"
	"namecheap-cli/utils"
	"os"
	"path/filepath"
	"strings"
)

// variables
var (
	emailAddresses    []string
	domainNames       []string
	outputPath        string
	defaultOutputPath string

	keyFilename   string
	csrFilename   string
	domainName    string
	keyOutputPath string
	csrOutputPath string

	defaultEmails  = []string{consts.EmailAddress}
	defaultDomains = []string{}
)

var SSLGenerateCSRCmd = &cobra.Command{
	Use:   "generate-csr",
	Short: "Create a new Certificate Signing Request",
	Long: `
	Create a new certificate signing request

	Valid values (case-insensitive):
	--domain	= <root domain or wildcard domain> 
	--email		= <email address> (DEFAULT: value set in const.go)

	Example:
	namecheap ssl generate-csr --domain <domain> --email <email address for the cert>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// save private key

		if strings.Contains(domainNames[0], "*") {
			domainName = strings.ReplaceAll(domainNames[0], "*", "wildcard")
			keyFilename = domainName + ".key"
			csrFilename = domainName + ".csr"
		} else {
			keyFilename = domainNames[0] + ".key"
			csrFilename = domainNames[0] + ".csr"
		}

		if outputPath == "" {
			keyOutputPath = filepath.Join(defaultOutputPath, keyFilename)
			csrOutputPath = filepath.Join(defaultOutputPath, csrFilename)
		} else {
			keyOutputPath = filepath.Join(outputPath, keyFilename)
			csrOutputPath = filepath.Join(outputPath, csrFilename)
		}

		// generate private key
		privateKey := utils.GeneratePrivateKey()
		utils.SavePrivateKey(keyOutputPath, privateKey)
		csr := utils.GenerateCSR(domainNames, emailAddresses, privateKey)
		utils.SaveCSR(csrOutputPath, csr)
	},
}

func init() {
	// get current path
	defaultOutputPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	// set default output path
	defaultOutputPath = filepath.Join(defaultOutputPath, "csr")

	// create the directory if not exist
	err = os.MkdirAll(defaultOutputPath, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	SSLCmd.AddCommand(SSLGenerateCSRCmd)
	// add flags
	SSLGenerateCSRCmd.Flags().StringArrayVarP(&domainNames, "domain", "d", defaultDomains, "Domain name for ceritifcate to be issued for.")
	SSLGenerateCSRCmd.Flags().StringArrayVarP(&emailAddresses, "email", "e", defaultEmails, "Email addresses for certificate to be issued for.")
	SSLGenerateCSRCmd.Flags().StringVarP(&outputPath, "output", "o", defaultOutputPath, "Output path for the generated CSR and private key.")
	SSLGenerateCSRCmd.MarkFlagRequired("domain")
	SSLGenerateCSRCmd.MarkFlagRequired("email")
}
