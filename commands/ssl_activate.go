package commands

import (
	"github.com/spf13/cobra"
	"namecheap-cli/common"
	"namecheap-cli/consts"
)

var (
	certID       int
	csrFile      string
	emailAddress string
)

var SSLActivateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Initiates creation of a new certifcate",
	Long: `
	Initiates creation of a new certifcate

	Valid values (case-insensitive):
	--certid	= <certid in certificates list> 
	--csr		= <path to csr file> (DEFAULT: value set in const.go)
	--email		= <email address> (DEFAULT: email address set in const)

	Example:
	namecheap ssl activate --certid <certid> --csr <path to csr file> --email <email address for the cert>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// save private key
		common.SSLActivate(certID, csrFile)
	},
}

func init() {
	SSLCmd.AddCommand(SSLActivateCmd)
	//// add flags
	SSLActivateCmd.Flags().IntVarP(&certID, "certid", "i", 0, "Unique identifier of SSL certificate to be issued.")
	SSLActivateCmd.Flags().StringVarP(&csrFile, "csr", "c", "", "Path to CSR file.")
	SSLActivateCmd.Flags().StringVarP(&emailAddress, "email", "e", consts.EmailAddress, "Email address to send signed SSL certificate file to.")
	SSLActivateCmd.MarkFlagRequired("certid")
	SSLActivateCmd.MarkFlagRequired("csr")
}
