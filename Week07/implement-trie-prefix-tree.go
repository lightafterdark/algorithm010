package main

//208. 实现 Trie (前缀树)
type Trie struct {
	isWord bool
	next   map[string]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{next: make(map[string]*Trie), isWord: false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for _, w := range word {
		value, ok := root.next[string(w)]
		if ok {
			root = value
		} else {
			newNode := make(map[string]*Trie, 0)
			root.next[string(w)] = &Trie{
				next:   newNode,
				isWord: false,
			}
			root = root.next[string(w)]
		}
	}
	root.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	isWord, ok := this.searchWord(word)
	if ok && isWord {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	_, ok := this.searchWord(prefix)
	return ok
}

func (this *Trie) searchWord(word string) (bool, bool) {
	var (
		value *Trie
		ok    bool
	)
	root := this
	for _, v := range word {
		value, ok = root.next[string(v)]
		if ok {
			root = value
			continue
		}
		return false, false
	}
	return value.isWord, true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
