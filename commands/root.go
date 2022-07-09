package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "namecheap",
	Short: `
	A CLI tool built from Go
	that interact with NameCheap API`,
	SuggestionsMinimumDistance: 1,
}

func init() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}
