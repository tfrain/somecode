package structure

type Trie struct {
	childrens [26]*Trie
	isEnd     bool
}

// func Constructor() Trie {
// 	return Trie{}
// }

func (this *Trie) Insert(word string) {
	curr := this
	for _, w := range word {
		c := w-'a'
		if curr.childrens[c] == nil {
			curr.childrens[c] = &Trie{}
		}
		curr = curr.childrens[c]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, w := range word {
		c := w-'a'
		if curr.childrens[c] == nil {
			return false
		}
		curr = curr.childrens[c]
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, w := range prefix {
		c := w-'a'
		if curr.childrens[c] == nil {
			return false
		}
		curr = curr.childrens[c]
	}
	return curr.isEnd
}

/**
* Your Trie object will be instantiated and called as such:
* obj := Constructor();
* obj.Insert(word);
* param_2 := obj.Search(word);
* param_3 := obj.StartsWith(prefix);
 */
