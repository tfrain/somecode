package structure

type LRUNode struct {
	prev, nxt *LRUNode
	key, val  int
}

type LRUCache struct {
	head, tail *LRUNode
	m          map[int]*LRUNode
	c          int
}

// func Constructor(capacity int) LRUCache {
// 	head := &Node{
// 		key: 0,
// 		val: 0,
// 	}
// 	tail := &Node{
// 		key: 0,
// 		val: 0,
// 	}
// 	head.nxt = tail
// 	tail.prev = head
// 	return LRUCache{
// 		head: head,
// 		tail: tail,
// 		m:    make(map[int]*Node),
// 		c:    capacity,
// 	}
// }

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.remove(node)
		this.Insert(node)
		return node.key
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		this.remove(node)
	}
	if this.c == len(this.m) {
		this.remove(this.tail.prev)
	}
	this.Insert(&LRUNode{
		key: key,
		val: value,
	})
}

func (this *LRUCache) remove(node *LRUNode) {
	delete(this.m, node.key)
	node.prev.nxt = node.nxt
	node.nxt.prev = node.prev
}

func (this *LRUCache) Insert(node *LRUNode) {
	this.m[node.key] = node
	nxt := this.head.nxt
	this.head.nxt = node
	node.nxt = nxt
	nxt.prev = node
	node.prev = this.head
}
