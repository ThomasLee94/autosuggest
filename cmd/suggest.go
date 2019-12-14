package cmd

import (
	"fmt"
	"log"

	trie "github.com/ThomasLee94/autosuggest/trie"

	"github.com/ThomasLee94/autosuggest/save"
	"github.com/spf13/cobra"
)

//cmd to add auto-suggestions
var suggestCmd = &cobra.Command{
	Use: "	",
	Short: "show all added suggests!",
	Run: func(cmd *cobra.Command, args []string) {
		var trieObj trie.Trie
		// load tri obj back
		if err := save.Load("./file.tmp", trieObj); err != nil {
			log.Fatalln(err)
		}

		// var complete []string

		// for i := range args {
		// 	complete = append(complete, args[i])
		// }

		fmt.Printf("%v", args)
		// complete := trieObj.Complete(args)
		// fmt.Printf("%v", complete)

	},
}

//	Adds show command to root command
func init() {
	RootCmd.AddCommand(suggestCmd)
}
