package structure

import "math/rand"

type RandomizedSet struct {
	nums       []int
	valToIndex map[int]int
}

// func Constructor() RandomizedSet {
// 	return RandomizedSet{
// 		nums:       make([]int, 0),
// 		valToIndex: make(map[int]int),
// 	}
// }

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToIndex[val]; ok {
		return false
	}
	// 放到尾部
	this.valToIndex[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.valToIndex[val]; !ok {
		return false
	}
	index := this.valToIndex[val]
	// 最后一个元素索引值改成index
	this.valToIndex[this.nums[len(this.nums)-1]] = index
	// 交换nums中index和最后一位的数值
	this.nums[index], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[index]
	// 数组中删除元素
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.valToIndex, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
