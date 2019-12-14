package cmd

import (
	"bufio"
	"os"

	"github.com/ThomasLee94/autosuggest/trie"
	"github.com/spf13/cobra"
)

//cmd to add auto-suggestions
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add commonly used commands for auto-suggestions!",
	Run: func(cmd *cobra.Command, args []string, trieObj trie.Trie) {
		// load tri obj back
		// var trieObj trie.Trie
		// if err := save.Load("./file.tmp", trieObj); err != nil {
		// 	log.Fatalln(err)
		// }

		// fmt.Println(trieObj)

		// create reader to read from standard input
		reader := bufio.NewReader(os.Stdin)

		// saves chars from terminal input
		chars, _ := reader.ReadString('\n')
		// insert strings into trie
		trieObj.Insert(string(chars))

		// if err := save.Save("./file.tmp", trieObj); err != nil {
		// 	log.Fatalln(err)
		// }

	},
}

//	Adds show command to root command
// func init() {
// 	RootCmd.AddCommand(addCmd)
// }
