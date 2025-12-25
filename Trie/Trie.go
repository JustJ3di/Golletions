package trie

type trienode struct {
	children map[rune]*trienode
	end      bool
}

type Trie struct {
	root *trienode
}

func New() *Trie {
	return &Trie{
		root: &trienode{children: make(map[rune]*trienode)},
	}
}

func (t *Trie) Search(key string) bool {

	curr := t.root
	for _, ch := range key {
		if _, exist := curr.children[ch]; !exist {
			return false
		}
		curr = curr.children[ch]
	}
	return curr.end
}

func (t *Trie) Insert(str string) {
	curr := t.root
	for _, ch := range str {
		if _, exist := curr.children[ch]; !exist {
			curr.children[ch] = &trienode{children: make(map[rune]*trienode)}
		}
		curr = curr.children[ch]
	}
	curr.end = true
}

// Find only the prefix it return true if the prefix is in the trie, not if the last rune in prefix is the end of the world
func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for _, ch := range prefix {
		if _, exist := curr.children[ch]; !exist {
			return false
		}
		curr = curr.children[ch]
	}
	return true
}
