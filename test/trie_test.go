package test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/ThomasLee94/autosuggest/trie"
)

/* -------------------------------------------------------------------------- */
/*                               TRIE INIT TEST                               */
/* -------------------------------------------------------------------------- */

func TestTrieAttributes(t *testing.T) {
	// init trie obj
	trieObj := trie.NewTrie()

	// test size
	assert.Equal(t, trieObj.Size, 1)

	// test root attributes
	assert.Equal(t, trieObj.Root.Character, "")
	assert.Equal(t, trieObj.Root.IsTerminal(), false)
	assert.Equal(t, len(trieObj.Root.Children), 0)

}

/* -------------------------------------------------------------------------- */
/*                         TRIE INIT WITH STRING TEST                         */
/* -------------------------------------------------------------------------- */

func TestTrieWithString(t *testing.T) {
	// init trie obj
	trieObj := trie.NewTrie("A")

	// test root attributes
	assert.Equal(t, trieObj.Root.Character, "A")
	assert.Equal(t, trieObj.Root.IsTerminal(), false)
	assert.Equal(t, len(trieObj.Root.Children), 0)

}

/* -------------------------------------------------------------------------- */
/*                              TRIE INSERT TEST                              */
/* -------------------------------------------------------------------------- */

func TestInsert(t *testing.T) {
	// init trie obj
	trieObj := trie.NewTrie()

	prefix := "AB"
	trieObj.Insert(prefix)

	// test root attributes
	assert.Equal(t, trieObj.Root.Character, "")
	assert.Equal(t, trieObj.Root.IsTerminal(), false)
	assert.Equal(t, len(trieObj.Root.Children), 1)

	// test 'A' node
	childNodeA, _ := trieObj.Root.GetChildren("A")
	assert.Equal(t, childNodeA.Character, "")
	assert.Equal(t, childNodeA.IsTerminal(), false)
	assert.Equal(t, len(childNodeA.Children), 1)
	assert.Equal(t, childNodeA.HasChildren("B"), true)

	// test 'B' node
	childNodeB, _ := childNodeA.GetChildren("B")
	assert.Equal(t, childNodeB.Character, "B")
	assert.Equal(t, childNodeB.IsTerminal(), true)
	assert.Equal(t, len(childNodeB.Children), 0)

}

/* -------------------------------------------------------------------------- */
/*                          TRIE INSERT MULTIPLE TEST                         */
/* -------------------------------------------------------------------------- */

func TestMultipleInsert(t *testing.T) {
	// TODO: FIX NODE CHARACTER KEY TESTS

	// prefixes for insertion
	prefix1 := "ABC"
	prefix2 := "ABE"
	prefix3 := "A"
	prefix4 := "EFG"

	// init trie obj
	trieObj := trie.NewTrie()

	// ─── INSERT PREFIX1 ───────────────────────────────────────────────────────
	trieObj.Insert(prefix1)

	// test node A
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

	// test node B
	childNodeB, _ := childNodeA.GetChildren("B")

	if childNodeB.Character != "B" {
		t.Errorf("Node does not have a child containing prefex %s", "B")
	}

	if childNodeB.IsTerminal() {
		t.Errorf("Node of trie is %t; want %t", childNodeA.IsTerminal(), false)
	}

	if len(childNodeB.Children) != 1 {
		t.Errorf("node B has %d children; want %d", childNodeB.Children, 1)
	}

	if !childNodeB.HasChildren("C") {
		t.Errorf("node B has %d children; want %d", childNodeB.Children, 1)
	}

	// test node C
	childNodeC, _ := childNodeA.GetChildren("C")

	if childNodeC.Character != "C" {
		t.Errorf("Node does not have a child containing prefex %s", "C")
	}

	if childNodeC.IsTerminal() {
		t.Errorf("Node of trie is %t; want %t", childNodeA.IsTerminal(), true)
	}

	if len(childNodeC.Children) != 1 {
		t.Errorf("node A has %d children; want %d", childNodeB.Children, 0)
	}

	// ─── INSERT PREFIX2 ───────────────────────────────────────────────────────
	trieObj.Insert(prefix2)

	// test node a
	lenOfChildren := len(childNodeA.Children)
	if lenOfChildren != 1 {
		t.Errorf("Root has %d number of children, want %d", lenOfChildren, 1)
	}

	// test node A
	if childNodeA.Character != "" {
		t.Errorf("Current node should have character %s", "A")
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

	// test node b
	if childNodeA.Character != "" {
		t.Errorf("Current node should have character %s", "A")
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

	assert.Equal(t, lenOfChildren, 1)

}
