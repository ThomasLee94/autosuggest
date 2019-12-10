package cmd

import (
	"bufio"
	"os"

	"github.com/ThomasLee94/autosuggest/trie"
	"github.com/spf13/cobra"
)

//RootCmd for autosuggest
var suggestCmd = &cobra.Command{
	Use:   "add",
	Short: "Listening to all commands for auto-suggestions!",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: concurrently insert user input & showing completions text
		trieObj := trie.NewTrie()
		// trie object channel to allow for grayed out text
		stringChannel := make(chan string)
		// create reader to read from standard input
		reader := bufio.NewReader(os.Stdin)

		// goroutine to build trie with user input being streamed in
		// stream characters as it is being typed
		func()  {
			// saves chars from terminal input
			chars, _ := reader.ReadString('\n')
			// insert strings into trie
			for char := range chars {
				trieObj.Insert(string(char))
				stringChannel <- string(char)
			}
		}()

		// goroutine to show grayed out text
		func() []string{
			return trieObj.Complete(<-stringChannel)
		}()

		//input := <-trieChannel
	},
}

//	Adds show command to root command
func init() {
	RootCmd.AddCommand(suggestCmd)
}
