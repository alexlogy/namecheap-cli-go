package commands

import (
	"github.com/spf13/cobra"
	"namecheap-cli/common"
)

var SSLListCmd = &cobra.Command{
	Use:   "list",
	Short: "Returns a list of SSL certificates for a particular user",
	Long: `
	Returns a list of SSL certificates for a particular user
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// process ssl type
		common.SSLList()
	},
}

func init() {
	SSLCmd.AddCommand(SSLListCmd)
}
