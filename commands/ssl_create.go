package commands

import (
	"github.com/spf13/cobra"
	"log"
	"namecheap-cli/common"
	"strings"
)

// variables
var (
	sslType = map[string]string{
		"single":   "PositiveSSL",
		"wildcard": "PositiveSSL Wildcard",
	}
	certType string
	years    int
)

var SSLCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new SSL certificate",
	Long: `
	Create an order of a SSL certificate according to type
	and years.

	Valid values (case-insensitive):
	--type	= single, wildcard (DEFAULT: single)
	--years	= 1, 2, 3, 4, 5 (DEFAULT: 1)

	Example:
	namecheap ssl create
	namecheap ssl create --type wildcard --years 2
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// process ssl type
		if sslType[strings.ToLower(certType)] == "" {
			log.Fatalln("[error] invalid ssl certificate type! allowed: single or wildcard")
		}

		if years < 1 && years > 5 {
			log.Fatalln("[error] invalid number of years to be issued! allowed: 1 to 5")
		}

		// validation succeeded
		common.SSLCreate(sslType[strings.ToLower(certType)], years)
	},
}

func init() {
	SSLCmd.AddCommand(SSLCreateCmd)
	// add flags
	SSLCreateCmd.Flags().StringVarP(&certType, "type", "t", "single", "SSL Certificate Type (single/wildcard). Default to single.")
	SSLCreateCmd.Flags().IntVarP(&years, "years", "y", 1, "No. of years SSL issued for. Default to 1.")
}
