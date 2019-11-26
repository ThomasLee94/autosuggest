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

	assert.Equal(t, childNodeA.Character, "A")
	assert.Equal(t, childNodeA.IsTerminal(), false)
	assert.Equal(t, len(childNodeA.Children), 1)
	assert.Equal(t, childNodeA.HasChildren("B"), true)

	// test node B
	childNodeB, _ := childNodeA.GetChildren("B")

	assert.Equal(t, childNodeB.Character, "B")
	assert.Equal(t, childNodeB.IsTerminal(), false)
	assert.Equal(t, len(childNodeB.Children), 1)
	assert.Equal(t, childNodeB.HasChildren("C"), true)

	// test node C
	childNodeC, _ := childNodeA.GetChildren("C")

	assert.Equal(t, childNodeC.Character, "C")
	assert.Equal(t, childNodeC.IsTerminal(), true)
	assert.Equal(t, len(childNodeC.Children), 0)

	// ─── INSERT PREFIX2 ───────────────────────────────────────────────────────
	trieObj.Insert(prefix2)

	// test overlap
	assert.Equal(t, len(childNodeA.Children), 1)

	// test node A
	assert.Equal(t, childNodeA.Character, "A")
	assert.Equal(t, childNodeA.IsTerminal(), false)
	assert.Equal(t, len(childNodeA.Children), 1)
	assert.Equal(t, childNodeA.HasChildren("B"), true)

	// test node b
	assert.Equal(t, childNodeB.Character, "B")
	assert.Equal(t, childNodeB.IsTerminal(), false)
	assert.Equal(t, len(childNodeB.Children), 2)
	assert.Equal(t, childNodeB.HasChildren("B"), true)
	assert.Equal(t, childNodeB.HasChildren("C"), true)

}
