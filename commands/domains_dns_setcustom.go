package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"namecheap-cli/common"
)

var DomainDNSSetCustomCmd = &cobra.Command{
	Use:   "set-custom",
	Short: "set custom nameservers for a specified domain",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 3 {
			nameServers := []string{args[1], args[2]}
			common.DomainsDNSSetCustom(args[0], nameServers)
		} else {
			fmt.Println("[error] invalid number of arguments")
		}
	},
}

func init() {
	DomainDNSCmd.AddCommand(DomainDNSSetCustomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
