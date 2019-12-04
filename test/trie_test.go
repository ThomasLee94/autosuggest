package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ThomasLee94/autosuggest/trie"
)

/* -------------------------------------------------------------------------- */
/*                               TRIE INIT TEST                               */
/* -------------------------------------------------------------------------- */

func TestTrieAttributes(t *testing.T) {
	// init trie obj
	trieObj := trie.NewTrie()

	// test size
	assert.Equal(t, trieObj.Size, 0)

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
	assert.Equal(t, trieObj.Root.Character, "")
	assert.Equal(t, trieObj.Root.IsTerminal(), false)
	assert.Equal(t, len(trieObj.Root.Children), 1)
	assert.Equal(t, trieObj.Root.HasChildren("A"), true)

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

	// test "A" node
	childNodeA, _ := trieObj.Root.GetChildren("A")
	assert.Equal(t, childNodeA.Character, "")
	assert.Equal(t, childNodeA.IsTerminal(), false)
	assert.Equal(t, len(childNodeA.Children), 1)
	assert.Equal(t, childNodeA.HasChildren("B"), true)

	// test "B" node
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
	// prefix "ABC"
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
	// prefix "ABE"
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

	// test new node E
	childNodeE, _ := childNodeB.GetChildren("E")

	assert.Equal(t, childNodeE.Character, "E")
	assert.Equal(t, childNodeE.IsTerminal(), true)
	assert.Equal(t, len(childNodeE.Children), 0)

	// ─── INSERT PREFIX3 ───────────────────────────────────────────────────────
	// prefix "A"
	trieObj.Insert(prefix3)

	// test root node
	assert.Equal(t, trieObj.Root.Character, "")
	assert.Equal(t, trieObj.Root.IsTerminal(), false)
	assert.Equal(t, len(trieObj.Root.Children), 1)
	assert.Equal(t, trieObj.Root.HasChildren("A"), true)

	// test node A
	assert.Equal(t, childNodeA.Character, "A")
	assert.Equal(t, childNodeA.IsTerminal(), false)
	assert.Equal(t, len(childNodeA.Children), 1)
	assert.Equal(t, childNodeA.HasChildren("B"), true)

	// ─── INSERT PREFIX4 ───────────────────────────────────────────────────────
	// prefix "EFG"
	trieObj.Insert(prefix4)

	// test root node
	assert.Equal(t, trieObj.Root.Character, "")
	assert.Equal(t, trieObj.Root.IsTerminal(), false)
	assert.Equal(t, len(trieObj.Root.Children), 2)
	assert.Equal(t, trieObj.Root.HasChildren("A"), true)
	assert.Equal(t, trieObj.Root.HasChildren("E"), true)

	// test node E
	rootChildNodeE, _ := trieObj.Root.GetChildren("E")
	assert.Equal(t, rootChildNodeE.Character, "E")
	assert.Equal(t, rootChildNodeE.IsTerminal(), false)
	assert.Equal(t, len(rootChildNodeE.Children), 1)
	assert.Equal(t, rootChildNodeE.HasChildren("F"), true)

	// test node F
	childNodeF, _ := rootChildNodeE.GetChildren("F")
	assert.Equal(t, childNodeF.Character, "F")
	assert.Equal(t, childNodeF.IsTerminal(), false)
	assert.Equal(t, len(childNodeF.Children), 1)
	assert.Equal(t, childNodeF.HasChildren("G"), true)

	// test node G
	childNodeG, _ := rootChildNodeE.GetChildren("G")
	assert.Equal(t, childNodeG.Character, "G")
	assert.Equal(t, childNodeG.IsTerminal(), true)
	assert.Equal(t, len(childNodeG.Children), 0)

}

/* -------------------------------------------------------------------------- */
/*                           TEST SIZE AND IS EMPTY                           */
/* -------------------------------------------------------------------------- */

func TestSize(t *testing.T) {
	// init trie obj
	trieObj := trie.NewTrie()

	// test init
	assert.Equal(t, trieObj.Size, 0)
	assert.Equal(t, trieObj.IsEmpty(), true)

	// test size with 1st insert
	trieObj.Insert("A")
	assert.Equal(t, trieObj.Size, 1)
	assert.Equal(t, trieObj.IsEmpty(), false)

	// test size with 2nd insert
	trieObj.Insert("ABC")
	assert.Equal(t, trieObj.Size, 2)
	assert.Equal(t, trieObj.IsEmpty(), false)

	// test size with 3rd insert
	trieObj.Insert("ABE")
	assert.Equal(t, trieObj.Size, 3)
	assert.Equal(t, trieObj.IsEmpty(), false)

	// test after 4th insert
	trieObj.Insert("HIJ")
	assert.Equal(t, trieObj.Size, 4)
	assert.Equal(t, trieObj.IsEmpty(), false)

}

/* -------------------------------------------------------------------------- */
/*                              TEST SIZE REPEAT                              */
/* -------------------------------------------------------------------------- */

func TestSizeRepeat(t *testing.T) {
	// init trie obj
	trieObj := trie.NewTrie()

	// test size with 1st insert
	trieObj.Insert("A")
	assert.Equal(t, trieObj.Size, 1)
	assert.Equal(t, trieObj.IsEmpty(), false)

	// test size with repeated insert
	trieObj.Insert("A")
	assert.Equal(t, trieObj.Size, 1)

	// test size with 2nd insert
	trieObj.Insert("ALO")
	assert.Equal(t, trieObj.Size, 2)
	assert.Equal(t, trieObj.IsEmpty(), false)

	// test with repeated insert
	trieObj.Insert("ALO")
	assert.Equal(t, trieObj.Size, 2)

	// test with 3rd insert
	trieObj.Insert("ALOHA")
	assert.Equal(t, trieObj.Size, 3)
	assert.Equal(t, trieObj.IsEmpty(), false)

	// test with repeated insert
	trieObj.Insert("ALOHA")
	assert.Equal(t, trieObj.Size, 3)

}

/* -------------------------------------------------------------------------- */
/*                                TEST CONTAINS                               */
/* -------------------------------------------------------------------------- */

func TestContains(t *testing.T) {
	// init trie obj
	trieObj := trie.NewTrie("ABC", "ABD", "A", "XYZ")

	// Test contains for all substrings
	assert.Equal(t, trieObj.Contains("ABC"), true)
	assert.Equal(t, trieObj.Contains("ABD"), true)
	assert.Equal(t, trieObj.Contains("AB"), false)
	assert.Equal(t, trieObj.Contains("BC"), false)
	assert.Equal(t, trieObj.Contains("BD"), false)
	assert.Equal(t, trieObj.Contains("A"), true)
	assert.Equal(t, trieObj.Contains("B"), false)
	assert.Equal(t, trieObj.Contains("C"), false)
	assert.Equal(t, trieObj.Contains("D"), false)
	assert.Equal(t, trieObj.Contains("XYZ"), true)
	assert.Equal(t, trieObj.Contains("XY"), false)
	assert.Equal(t, trieObj.Contains("YZ"), false)
	assert.Equal(t, trieObj.Contains("X"), false)
	assert.Equal(t, trieObj.Contains("Y"), false)
	assert.Equal(t, trieObj.Contains("Z"), false)

}

/* -------------------------------------------------------------------------- */
/*                                TEST COMPLETE                               */
/* -------------------------------------------------------------------------- */

func TestComplete(t *testing.T) {
	// init trie obj
	trieObj := trie.NewTrie("ABC", "ABD", "A", "XYZ")

	var emptySlice []string

	// test complete
	assert.Equal(t, trieObj.Complete("ABC"), [1]string{"ABC"})
	assert.Equal(t, trieObj.Complete("ABD"), [1]string{"ABD"})
	assert.Equal(t, trieObj.Complete("AB"), [2]string{"ABC", "ABD"})
	assert.Equal(t, trieObj.Complete("BC"), emptySlice)
	assert.Equal(t, trieObj.Complete("BD"), emptySlice)
	assert.Equal(t, trieObj.Complete("A"), [3]string{"A", "ABC", "ABD"})
	assert.Equal(t, trieObj.Complete("B"), emptySlice)
	assert.Equal(t, trieObj.Complete("C"), emptySlice)
	assert.Equal(t, trieObj.Complete("D"), emptySlice)
	assert.Equal(t, trieObj.Complete("XYZ"), [1]string{"XYZ"})
	assert.Equal(t, trieObj.Complete("XY"), [1]string{"XYZ"})
	assert.Equal(t, trieObj.Complete("YZ"), emptySlice)
	assert.Equal(t, trieObj.Complete("X"), [1]string{"XYZ"})
	assert.Equal(t, trieObj.Complete("Y"), emptySlice)
	assert.Equal(t, trieObj.Complete("Z"), emptySlice)

}

func TestStrings(t *testing.T) {
	// init trie obj
	trieObj := trie.NewTrie()

	insertStrings := [4]string{"ABC", "ABD", "A", "JKL"}
	var outputStrings []string

	for _, word := range insertStrings {
		trieObj.Insert(word)
		outputStrings = append(outputStrings, word)

		// test tree can retrieve all strings that have been inserted
		trieStrings := trieObj.Strings()
		assert.Equal(t, len(outputStrings), len(trieStrings))
		assert.ElementsMatch(t, outputStrings, trieStrings)
	}

}
