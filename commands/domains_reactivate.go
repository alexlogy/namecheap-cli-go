package commands

import (
	"namecheap-cli/common"

	"github.com/spf13/cobra"
)

var DomainReactivateCmd = &cobra.Command{
	Use:   "reactivate",
	Short: "reactivate domain",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		common.DomainsReactivate(args[0])
	},
}

func init() {
	DomainsCmd.AddCommand(DomainReactivateCmd)
}
