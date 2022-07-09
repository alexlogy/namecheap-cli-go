package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"namecheap-cli/common"
)

var DomainCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "check availability for a specified domain",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 && len(args) < 50 {
			common.DomainsCheck(args)
		} else {
			fmt.Println("[error] invalid number of arguments")
		}
	},
}

func init() {
	DomainsCmd.AddCommand(DomainCheckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
