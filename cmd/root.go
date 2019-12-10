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
		// TODO: concurrently insert user input & showing completions text

		trie := trie.NewTrie()
		// trie object channel to allow for grayed out text
		trieChannel := make(chan trie.Trie)

		// goroutine to build trie with user input being streamed in
		go func(trieChannel chan) {
			// create reader to read from standard input
			reader := bufio.NewReader(os.Stdin)
			// saves chars from terminal input
			chars, _ := reader.ReadString('\n')
			// insert strings into trie
			for char := range chars {
				trie.Insert(string(char))
			}

			trieChannel <- trie
		}(trieChannel)

		// goroutine to show grayed out text
		go func(trieChannel chan) []string{
			return trieChannel.trie.Complete()
		}(trieChannel)

		input := <- trieChannel
	},
}

//	Adds show command to root command
func init() {
	RootCmd.AddCommand(RootCmd)
}
