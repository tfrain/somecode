package structure

type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*LRUNode
	head     *LRUNode
	tail     *LRUNode
}

type LRUNode struct {
	key   int
	value int
	prev  *LRUNode
	next  *LRUNode
}

// func Constructor(capacity int) LRUCache {
// 	head, tail := new(LRUNode), new(LRUNode)
// 	return LRUCache{
// 		capacity: capacity,
// 		size:     0,
// 		cache:    make(map[int]*LRUNode),
// 		head:     head,
// 		tail:     tail,
// 	}
// }

func (this *LRUCache) Get(key int) int {
	if LRUNode, exists := this.cache[key]; exists {
		this.moveToHead(LRUNode)
		return LRUNode.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.value = value
		this.moveToHead(node)
	} else {
		newNode := &LRUNode{
			key:   key,
			value: value,
		}
		this.cache[key] = newNode
		this.addToHead(newNode)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	}
}

func (this *LRUCache) addToHead(node *LRUNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *LRUNode {
	node := this.tail.prev
	node.prev.next = node.next
	node.next.prev = node.prev
	return node
}

// 这种会频繁更新 map, 导致创建和销毁，map 视图应该减少更新
// func (this *LRUCache) insert(n *Node) {
// 	this.m[n.key] = n
// 	next := this.head.Next
// 	this.head.Next = n
// 	// 这个能提前
// 	n.Next = next
// 	n.Prev = this.head
// 	next.Prev = n

// }

// func (this *LRUCache) remove(n *Node) {
// 	delete(this.m, n.key)
// 	n.Prev.Next = n.Next
// 	n.Next.Prev = n.Prev
// }
