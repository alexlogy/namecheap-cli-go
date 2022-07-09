package main

import (
	"namecheap-cli/commands"
	"os"
)

func main() {
	err := commands.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
