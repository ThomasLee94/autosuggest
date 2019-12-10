package cmd

import (
	"bufio"
	"os"

	"github.com/ThomasLee94/autosuggest/trie"
	"github.com/spf13/cobra"
)

//cmd to add auto-suggestions
var suggestCmd = &cobra.Command{
	Use:   "add",
	Short: "add commonly used commands for auto-suggestions!",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: concurrently insert user input & showing completions text
		trieObj := trie.NewTrie()
		// trie object channel to allow for grayed out text
		stringChannel := make(chan string)
		// create reader to read from standard input
		reader := bufio.NewReader(os.Stdin)

		// saves chars from terminal input
		chars, _ := reader.ReadString('\n')
		// insert strings into trie
		for char := range chars {
			trieObj.Insert(string(char))
			stringChannel <- string(char)
		}
	},
}

//	Adds show command to root command
func init() {
	RootCmd.AddCommand(suggestCmd)
}
