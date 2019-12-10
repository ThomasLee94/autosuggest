package cmd

import (
	"github.com/spf13/cobra"
)

//cmd to add auto-suggestions
var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "show all added suggests!",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

//	Adds show command to root command
func init() {
	RootCmd.AddCommand(suggestCmd)
}
