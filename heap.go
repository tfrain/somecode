package leetcode

import "container/heap"

type MedianFinder struct {
	s, b *Heap
}

func Constructor() MedianFinder {
	s := &Heap{
		less: func(i, j int) bool {
			return i < j
		},
	}
	b := &Heap{
		less: func(i, j int) bool {
			return i > j
		},
	}
	return MedianFinder{
		s: s,
		b: b,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if (this.s.Len()+this.b.Len())%2 == 0 {
		heap.Push(this.s, num)
		heap.Push(this.b, heap.Pop(this.s))
	} else {
		heap.Push(this.b, num)
		heap.Push(this.s, heap.Pop(this.b))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.s.Len()+this.b.Len())%2 == 0 {
		return float64(this.b.h[0]+this.s.h[0]) / 2
	} else {
		return float64(this.b.h[0])
	}
}

type Less func(i, j int) bool

type Heap struct {
	less Less
	h    []int
}

func (h Heap) Swap(i, j int) {
	h.h[i], h.h[j] = h.h[j], h.h[i]
}

func (h Heap) Len() int {
	return len(h.h)
}

func (h Heap) Less(i, j int) bool {
	return h.less(h.h[i], h.h[j])
}

func (h *Heap) Push(x interface{}) {
	h.h = append(h.h, x.(int))
}

func (h *Heap) Pop() (x interface{}) {
	x, h.h = h.h[len(h.h)-1], h.h[:len(h.h)-1]
	return
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PQ, 0)
	for _, node := range lists {
		if node != nil {
			pq = append(pq, node)
		}
	}
	if len(pq) == 0 {
		return nil
	}
	heap.Init(&pq)
	
	dummy := new(ListNode)
	list := dummy
	for len(pq) > 0 {
		n := heap.Pop(&pq)
		node := n.(*ListNode)
		list.Next = node
		list = list.Next

		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
	}
	return list
}

type PQ []*ListNode

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
}

func (pq PQ) Less(a, b int) bool {
	return pq[a].Val < pq[b].Val
}

func (pq *PQ) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PQ) Pop() (x interface{}) {
	x, *pq = (*pq)[len(*pq)-1], (*pq)[:len(*pq)-1]
	return
}
