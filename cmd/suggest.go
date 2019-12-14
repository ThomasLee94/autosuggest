package cmd

import (
	"fmt"
	"log"
	"strings"

	trie "github.com/ThomasLee94/autosuggest/trie"

	"github.com/ThomasLee94/autosuggest/save"
	"github.com/spf13/cobra"
)

//cmd to add auto-suggestions
var SuggestCmd = &cobra.Command{
	Use:   "s",
	Short: "show all added suggests!",
	Run: func(cmd *cobra.Command, args []string) {
		var trieObj trie.Trie
		// load tri obj back
		if err := save.Load("./file.tmp", trieObj); err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%v", args)
		slice := strings.Join(args, " ")
		complete := trieObj.Complete(slice)
		fmt.Printf("%v", complete)

	},
}

//	Adds show command to root command
func init() {
	RootCmd.AddCommand(suggestCmd)
}
