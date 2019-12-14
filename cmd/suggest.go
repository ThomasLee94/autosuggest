package cmd

import (
	"autosuggest/trie"
	"log"

	"github.com/ThomasLee94/autosuggest/save"
	"github.com/spf13/cobra"
)

//cmd to add auto-suggestions
var suggestCmd = &cobra.Command{
	Use: "	",
	Short: "show all added suggests!",
	Run: func(cmd *cobra.Command, args []string) {
		var obj trie.Trie
		// load tri obj back
		if err := save.Load("./file.tmp", obj); err != nil {
			log.Fatalln(err)
		}
	},
}

//	Adds show command to root command
func init() {
	RootCmd.AddCommand(suggestCmd)
}
