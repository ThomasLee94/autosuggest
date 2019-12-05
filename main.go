package main

import (
	"fmt"

	asdf "github.com/ThomasLee94/autosuggest/trie"
)

func main() {
	fmt.Println("ehl")

	trieObj := asdf.NewTrie("A")
	fmt.Println(trieObj.Root.Children)

}
