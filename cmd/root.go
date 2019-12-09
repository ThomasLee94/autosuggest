package cmd

import "github.com/spf13/cobra"

//RootCmd for autosuggest
var RootCmd = &cobra.Command{
	Use:   "",
	Short: "Listening to all commands for auto-suggestions!",
}