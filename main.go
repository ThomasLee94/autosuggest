package main

import (
	"log"

	"github.com/ThomasLee94/autosuggest/cmd"
	"github.com/ThomasLee94/autosuggest/save"
	"github.com/ThomasLee94/autosuggest/trie"
)

func main() {
	trie := trie.NewTrie()

	if err := save.Save("./file.tmp", trie); err != nil {
		log.Fatalln(err)
	}

	cmd.RootCmd.Execute()
}
