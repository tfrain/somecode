package leetcode

func findKthLargest(nums []int, k int) int {
	l, r, pos := 0, len(nums)-1, len(nums)-k
	for l < r {
		idx := quickSortPartiion(nums, l, r)
		if idx > pos {
			r = idx - 1
		} else if idx < pos {
			l = idx + 1
		} else {
			return nums[idx]
		}
	}
	return -1
}

// 从小到大
func quickSortPartiion(nums []int, i, j int) int {
	l, r := i+1, j
	for l <= r {
		for l <= r && nums[l] <= nums[i] {
			l++
		}
		for l <= r && nums[r] > nums[i] {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[r], nums[i] = nums[i], nums[r]
	return r
}

func quickSortRecursive(arr []int, left, right int) {
	if left < right {
		pivotIndex := partition2(arr, left, right)
		quickSortRecursive(arr, left, pivotIndex-1)
		quickSortRecursive(arr, pivotIndex+1, right)
	}
}

func quickSortIterative(arr []int, left, right int) {
	stack := make([]int, right-left+1)
	top := -1

	top++
	stack[top] = left
	top++
	stack[top] = right

	for top >= 0 {
		right = stack[top]
		top--
		left = stack[top]
		top--

		pivotIndex := partition2(arr, left, right)

		if pivotIndex-1 > left {
			top++
			stack[top] = left
			top++
			stack[top] = pivotIndex - 1
		}

		if pivotIndex+1 < right {
			top++
			stack[top] = pivotIndex + 1
			top++
			stack[top] = right
		}
	}
}

func partition2(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}
