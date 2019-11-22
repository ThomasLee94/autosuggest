package trie

import "github.com/ThomasLee94/autosuggest/node"

type Trie struct {
	root *node.Node
	size int
}
