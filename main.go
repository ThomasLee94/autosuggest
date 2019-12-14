package main

import (
	"log"

	"github.com/ThomasLee94/autosuggest/save"
	"github.com/ThomasLee94/autosuggest/trie"
	"github.com/spf13/cobra/cobra/cmd"
)

func main() {
	trie := trie.NewTrie()

	if err := save.Save("./file.tmp", trie); err != nil {
		log.Fatalln(err)
	}

	cmd.Execute()

}
