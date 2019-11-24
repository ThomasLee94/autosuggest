package test

import (
	"testing"

	"github.com/ThomasLee94/autosuggest/trie"
)

func TestTrieStructAttributes(t *testing.T) {
	// init node struct
	trieObj := trie.NewTrie()

	// test size
	if trieObj.Size != 0 {
		t.Errorf("Size of Trie is %d; want %d", trieObj.Size, 0)
	}

	// test root attributes
	if trieObj.Root.Character != "" {
		t.Errorf("Size of Trie is %s; want %s", trieObj.Root.Character, "")
	}

	if trieObj.Root.IsTerminal() {
		t.Errorf("Root node of trie is %t; want %t", trieObj.Root.IsTerminal(), false)
	}

	if len(trieObj.Root.Children) != 0 {
		t.Errorf("Root node has %d children; want %d", trieObj.Root.Children, 0)
	}

}
