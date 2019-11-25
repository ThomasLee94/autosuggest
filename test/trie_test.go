package test

import (
	"testing"

	"github.com/ThomasLee94/autosuggest/trie"
)

func TestTrieAttributes(t *testing.T) {
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

func TestTrieWithString(t *testing.T) {
	// init node struct
	trieObj := trie.NewTrie("A")

	// test root attributes
	if trieObj.Root.Character != "A" {
		t.Errorf("Size of Trie is %s; want %s", trieObj.Root.Character, "A")
	}

	if trieObj.Root.IsTerminal() {
		t.Errorf("Root node of trie is %t; want %t", trieObj.Root.IsTerminal(), false)
	}

	if len(trieObj.Root.Children) != 0 {
		t.Errorf("Root node has %d children; want %d", trieObj.Root.Children, 0)
	}
}

func TestInsert(t *testing.T) {
	// init node struct
	trieObj := trie.NewTrie()

	prefix := "AB"
	trieObj.Insert(prefix)

	// test root attributes
	if trieObj.Root.Character != "" {
		t.Errorf("Size of Trie is %s; want %s", trieObj.Root.Character, "")
	}

	if trieObj.Root.IsTerminal() {
		t.Errorf("Root node of trie is %t; want %t", trieObj.Root.IsTerminal(), false)
	}

	if len(trieObj.Root.Children) != 1 {
		t.Errorf("Root node has %d children; want %d", trieObj.Root.Children, 1)
	}

	// test 'A' node
	childNodeA, _ := trieObj.Root.GetChildren("A")
	if childNodeA.Character != "" {
		t.Errorf("Root node does not have a child containing prefex %s", "A")
	}

	if childNodeA.IsTerminal() {
		t.Errorf("Root node of trie is %t; want %t", childNodeA.IsTerminal(), false)
	}

	if len(childNodeA.Children) != 1 {
		t.Errorf("node A has %d children; want %d", childNodeA.Children, 1)
	}

	if !childNodeA.HasChildren("B") {
		t.Errorf("node A has %d children; want %d", childNodeA.Children, 1)
	}

	// test 'B' node
	childNodeB, _ := childNodeA.GetChildren("B")
	if childNodeB.Character != "B" {
		t.Errorf("node A does not have a child containing prefex %s", "B")
	}

	if !childNodeB.IsTerminal() {
		t.Errorf("Root node of trie is %t; want %t", childNodeA.IsTerminal(), true)
	}

	if len(childNodeB.Children) != 0 {
		t.Errorf("node A has %d children; want %d", len(childNodeB.Children), 0)
	}

}
