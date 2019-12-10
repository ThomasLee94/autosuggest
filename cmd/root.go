package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd for autosuggest
var RootCmd = &cobra.Command{
	Use:   "as",
	Short: "A CLI tool to add auto suggestions for your terminal!",
}

