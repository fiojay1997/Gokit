package engine

// node defines a trie node (contains patterns of the given url)
type node struct {
	pattern  string 
	part     string 
	children []*node 
	isWild   bool 
}