package cmd

import (
	"bufio"
	"os"

	"github.com/ThomasLee94/autosuggest/trie"
	"github.com/spf13/cobra"
)

//RootCmd for autosuggest
var RootCmd = &cobra.Command{
	Use:   "",
	Short: "Listening to all commands for auto-suggestions!",
	Run: func(cmd *cobra.Command, args []string) {
		// create reader to read from standard input
		reader := bufio.NewReader(os.Stdin)
		// saves chars from terminal input
		chars, _ := reader.ReadString('\n')
		// insert strings into trie
		trie := trie.NewTrie()
		for char := range chars {
			trie.Insert(string(char))
		}
	},
}

//	Adds show command to root command
func init() {
	RootCmd.AddCommand(RootCmd)
}
