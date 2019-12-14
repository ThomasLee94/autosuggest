package cmd

import (
	"bufio"
	"log"
	"os"

	"github.com/ThomasLee94/autosuggest/save"
	"github.com/ThomasLee94/autosuggest/trie"
	"github.com/spf13/cobra"
)

//cmd to add auto-suggestions
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add commonly used commands for auto-suggestions!",
	Run: func(cmd *cobra.Command, args []string) {
		// load tri obj back
		var trieObj trie.Trie
		if err := save.Load("./file.tmp", trieObj); err != nil {
			log.Fatalln(err)
		}
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
	RootCmd.AddCommand(addCmd)
}
