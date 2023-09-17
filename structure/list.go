package structure

type Node struct {
	nxt *Node
	val int
}

type MyLinkedList struct {
	head, tail *Node
	length     int
}

/** Initialize your data structure here. */
// func Constructor() MyLinkedList {
// 	head, tail := new(LRUNode), new(LRUNode)
// 	head.nxt = tail
// 	return MyLinkedList{
// 		head: head,
// 		tail: tail,
// 	}
// }

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}
	head := this.head
	for i := 0; i < index; i++ {
		head = head.nxt
	}
	return head.nxt.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	n := &Node{
		val: val,
	}
	nxt := this.head.nxt
	this.head.nxt = n
	n.nxt = nxt
	this.length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.tail.val = val
	this.tail.nxt = new(Node)
	this.tail = this.tail.nxt
	this.length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.length {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
	}
	if index == this.length {
		this.AddAtTail(val)
	}
	head := this.head
	for i := 0; i < index; i++ {
		head = head.nxt
	}
	n := &Node{
		val: val,
	}
	nxt := head.nxt
	head.nxt = n
	n.nxt = nxt
	this.length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		return
	}
	head := this.head
	for i := 0; i < index; i++ {
		head = head.nxt
	}
	head.nxt = head.nxt.nxt
	this.length--
}
