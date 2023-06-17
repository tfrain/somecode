package structure

import "math/rand"

// 384
type Solution struct {
	nums []int
}

// func Constructor(nums []int) Solution {
// 	return Solution{
// 		nums: nums,
// 	}
// }

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	random := make([]int, n)
	copy(random, this.nums)
	for i := range random {
		// 生成一个[i, n-1] 区间内的随机数
		// 对于 nums[0]，我们把它随机换到了索引 [0, n) 上，共有 n 种可能性；
		// 对于 nums[1]，我们把它随机换到了索引 [1, n) 上，共有 n - 1 种可能性；
		r := i + rand.Intn(n-i)
		swap(random, i, r)
	}
	return random
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

// 398
// type Solution struct {
// 	nums []int
// }

// func Constructor(nums []int) Solution {
// 	return Solution{
// 		nums: nums,
// 	}
// }

func (this *Solution) Pick(target int) int {
	var j int
	index := -1
	for i, v := range this.nums {
		if v == target {
			if rand.Intn(j+1) == 0 {
				index = i
			}
			j++
		}
	}
	return index
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 382

/* type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

func (this *Solution) GetRandom() int {
	head := this.head
	var j int
	var val int
	for head != nil {
		if 0 == rand.Intn(j+1) {
			val = head.Val
		}
		head = head.Next
		j++
	}
	return val
} */

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
