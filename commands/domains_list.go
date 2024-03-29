package commands

import (
	"fmt"
	"namecheap-cli/common"

	"github.com/spf13/cobra"
)

var (
	page    int
	listAll bool
)

var DomainListCmd = &cobra.Command{
	Use:   "list",
	Short: "list registered domains in namecheap account",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if listAll {
			common.DomainsGetListAll()
			return
		}

		if len(args) == 1 {
			common.DomainsGetList(args[0], page)
		} else if len(args) == 0 {
			common.DomainsGetList("", page)
		} else {
			fmt.Println("[error] invalid arguments")
		}

	},
}

func init() {
	DomainsCmd.AddCommand(DomainListCmd)

	// add flags
	DomainListCmd.Flags().IntVarP(&page, "page", "p", 1, "Page for Domains List. Default to Page 1.")

	DomainListCmd.Flags().BoolVarP(&listAll, "all", "a", false, "List all domains with each API call delayed to avoid throttling.")
}
