package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n-k+1)
	stack := []int{}
	for i, v := range nums {
		if len(stack) > 0 && stack[0] <= i-k+1 {
			stack = stack[1:]
		}
		for len(stack) > 0 && nums[stack[len(stack)-1]] < v {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		if i-k+1 >= 0 {
			res[i-k+1] = stack[0]
		}
	}
	return res
}

// 栈中元素不会额外排序(别想太多), 因为单调栈构造过程,就保证其单调
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := []int{}
	for i := n-1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[0]] {
			stack = stack[1:]
		}
		if len(stack) > 0 {
			res[i] = i-stack[0]
		} else {
			res[i] = 0
		}
		stack = append([]int{i}, stack...)
	}
	return res
}