package commands

import (
	"github.com/spf13/cobra"
)

var SSLCmd = &cobra.Command{
	Use:   "ssl",
	Short: "command to interact with ssl certificates in namecheap",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func init() {
	RootCmd.AddCommand(SSLCmd)
}
