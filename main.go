package main

import (
	"fmt"

	asdf "github.com/ThomasLee94/autosuggest/trie"
)

func main() {
	fmt.Println("ehl")

	trieObj := asdf.NewTrie()
	prefix1 := "ABC"
	prefix2 := "ABE"
	prefix3 := "A"
	prefix4 := "EFG"

	trieObj.Insert(prefix1)
	trieObj.Insert(prefix2)
	trieObj.Insert(prefix3)
	trieObj.Insert(prefix4)

	childNodeA, _ := trieObj.Root.GetChildren("A")
	childNodeB, _ := childNodeA.GetChildren("B")
	// childNodeE, _ := childNodeB.GetChildren("E")

}
