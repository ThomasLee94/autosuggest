package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd for autosuggest
var RootCmd = &cobra.Command{
	Use:   "a",
	Short: "A CLI tool to add 'auto suggestions' (not really though :P) for your terminal!",
}

