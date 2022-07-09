package commands

import (
	"github.com/spf13/cobra"
	"log"
	"namecheap-cli/common"
	"strconv"
)

var DomainRenewCmd = &cobra.Command{
	Use:   "renew <domain name> <no. of years to renew>",
	Short: "Renews an expiring domain",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// check whether 2nd argument is an int string
		years, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalln(err)
		}
		if years <= 9 {
			// renew the domain
			common.DomainsRenew(args[0], args[1])
		}
	},
}

func init() {
	DomainsCmd.AddCommand(DomainRenewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
