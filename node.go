package node

type Node struct {
	character string
	children  map[rune]*Node
	terminal  bool
}
