package trie

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	chars map[byte]*TrieNode
	size  int
}

func (tr *Trie) Count(word string) int {
	node := tr.root
	var ok bool
	var child *TrieNode
	for len(word) > 0 {
		if child, ok = node.chars[word[0]]; !ok {
			return 0
		}
		word = word[1:]
		node = child
	}
	return node.size
}

func (tr *Trie) Add(word string) {
	if len(word) == 0 {
		return
	}
	node := tr.root
	node.size++
	var ok bool
	var child *TrieNode
	for len(word) > 0 {
		if child, ok = node.chars[word[0]]; !ok {
			child = &TrieNode{chars: make(map[byte]*TrieNode)}
			node.chars[word[0]] = child
		}
		word = word[1:]
		node = child
		node.size++
	}
}

func (tr *Trie) AddPref(pref string) {
	size := tr.root.size
	newRoot := &TrieNode{}
	node := addPref(newRoot, []byte(pref), size)
	node.chars = tr.root.chars
	tr.root = newRoot
}

func addPref(node *TrieNode, pref []byte, size int) *TrieNode {
	node.size = size
	if len(pref) == 0 {
		return node
	}
	var child *TrieNode
	var ok bool
	if node.chars == nil {
		node.chars = make(map[byte]*TrieNode)
	}
	if child, ok = node.chars[pref[0]]; !ok {
		child = &TrieNode{}
		node.chars[pref[0]] = child
	}
	return addPref(child, pref[1:], size)
}
