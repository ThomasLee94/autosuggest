package cmd

import (
	"github.com/spf13/cobra"
	trie "github.com/ThomasLee94/autosuggest/trie"
)

trie := trie.NewTrie()

//RootCmd for autosuggest
var RootCmd = &cobra.Command{
	Use:   "a",
	Short: "A CLI tool to add 'auto suggestions' (not really though :P) for your terminal!",
}

// Execute executes the root command.
func Execute() error {
	return RootCmd.Execute()
}

func init() {

	RootCmd.AddCommand(AddCmd)
	RootCmd.AddCommand(SuggestCmd)
}
