package test

import (
	"testing"

	"github.com/ThomasLee94/autosuggest/trie"
)

func TestTrieStructAttributes(t *testing.T) {
	// init node struct
	trieObj := trie.NewTrie()

	// test size
	if trieObj.size != 0 {
		t.Errorf("Size of Trie is %d; want %d", trieObj.size, 0)
	}

	// test root attributes
	if trieObj.root.character != "" {
		t.Errorf("Size of Trie is %s; want %s", trieObj.root.character, "")
	}

	if trieObj.root.IsTerminal() {
		t.Errorf("Root node of trie is %t; want %t", trieObj.root.IsTerminal(), false)
	}

	if len(trieObje.root.Children) != 0 {
		t.Errorf("Root node has %d children; want %d", trieObje.root.Children, 0)
	}

}
