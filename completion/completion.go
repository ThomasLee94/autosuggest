// package completion will be used to pass slices as pointers
// in package trie for Complete(), Strings() & traverse.
package completion

type Completion struct {
	CompleteSlice []string
}