package structure

type mNode struct{
	val, min int
	next *mNode
}

type MinStack struct {
	head *mNode
}


// func Constructor() MinStack {
// 	return MinStack{}
// }


func (this *MinStack) Push(val int)  {
	min := val
	if this.head != nil && this.head.min < val {
		min = this.head.min
	}
	n := &mNode{
		val: val,
		min: min,
	}
	nxt := this.head
	this.head = n
	n.next = nxt
}


func (this *MinStack) Pop()  {
	this.head = this.head.next
}


func (this *MinStack) Top() int {
	return this.head.val
}


func (this *MinStack) GetMin() int {
	return this.head.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */