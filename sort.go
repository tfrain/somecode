package leetcode

import "fmt"

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

func QuickSort(a []int, left int, right int) {
	if a == nil || left < 0 || right <= 0 || left > right {
		return
	}
	var temp []int
	var i, j int
	temp = append(temp, right)
	temp = append(temp, left)
	for len(temp) > 0 {
		i = temp[len(temp)-1]
		temp = temp[:len(temp)-1]
		j = temp[len(temp)-1]
		temp = temp[:len(temp)-1]
		if i < j {
			k := Pritation(a, i, j)
			if k > i {
				temp = append(temp, k-1)
				temp = append(temp, i)
			}
			if j > k {
				temp = append(temp, j)
				temp = append(temp, k+1)
			}
		}
	}
}

func Pritation(a []int, left int, right int) int {
	pivot := a[right]
	i := left
	for j := left; j < right; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[right] = a[right], a[i]
	return i
}

func main() {
	arr := []int{3, 2, 1}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 324
func wiggleSort(nums []int) {
	// 1 是一个二进制数 0001。确保 n 总是奇数
	n := len(nums) | 1
	// 只去发现一次，发现 mid
	mid := findMiddle(nums)

	// ^1 是按位取反操作，将 1（二进制 0001）变成 -2（在二进制中，-2 表示为所有位都是 1 除了最低位是 0）。
	// &(^1) 是按位与操作，它将 len(nums)-1 的最低位清零，从而确保 lt 是一个偶数。
	// 这是因为在摆动排序中，我们希望小于中位数的元素位于偶数索引位置。
	gt, lt := 1, (len(nums)-1)&(^1)
	// 只要 i 是奇数，或者 i 没有超过 lt 的值
	// 更多像是防呆功能
	for i := 1; i&1 == 1 || i <= lt; {
		if i == n {
			i = 0
		}

		// 小于中位数的，放在偶数索引，这里是从最大的偶数索引下降
		if nums[i] < mid {
			nums[i], nums[lt] = nums[lt], nums[i]
			lt -= 2
		} else if nums[i] > mid {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt += 2
			i += 2
		} else {
			i += 2
		}
	}
}

func findMiddle(nums []int) int {
	lo, hi := 0, len(nums)-1
	mi := hi / 2
	for lo < hi {
		mid := newPartition(nums, lo, hi)
		if mi < mid {
			hi = mid - 1
		} else if mi > mid {
			lo = mid + 1
		} else {
			break
		}
	}
	// 这里 num 已经排好序了，从而将真正的中位数返回
	return nums[mi]
}

func newPartition(nums []int, lo, hi int) int {
	pivot := nums[lo]
	i, j := lo, hi+1
	for {
		for i+1 <= hi && nums[i+1] < pivot {
			i++
		}
		i++
		for nums[j-1] > pivot {
			j--
		}
		j--

		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}
