package interview

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main1() {
	rand.Seed(time.Now().UnixNano())
	nums := []int{rand.Intn(13) + 1, rand.Intn(13) + 1, rand.Intn(13) + 1, rand.Intn(13) + 1}
	fmt.Println(nums)
	if twentyFour(nums) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func twentyFour(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return nums[0] == 24
	}
	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}
			a, b := nums[i], nums[j]
			remaining := make([]int, 0, len(nums)-2)
			for k := range nums {
				if k != i && k != j {
					remaining = append(remaining, nums[k])
				}
			}
			if twentyFour(append(remaining, a+b)) ||
				twentyFour(append(remaining, a-b)) ||
				twentyFour(append(remaining, a*b)) ||
				(b != 0 && twentyFour(append(remaining, a/b))) {
				return true
			}
		}
	}
	return false
}

func judgePoint24(nums []int) bool {
	// 数字太大，要用 float64
	// nums 就相当于栈来使用
	var helper func([]float64) bool
	helper = func(nums []float64) bool {
		if len(nums) == 1 {
			return math.Abs(nums[0]-24) < 1e-6 // 0.000001
		}
		for i := range nums {
			for j := range nums {
				// 为了选择两个不同数字
				if i == j {
					continue
				}
				var next []float64
				for k, num := range nums {
					// 为了选择两个不同数字
					if k != i && k != j {
						next = append(next, num)
					}
				}
				// 将两个数字的每一种计算结果都加入列表计算
				for _, num := range getNums(nums[i], nums[j]) {
					if helper(append(next, num)) {
						return true
					}
				}
			}
		}
		return false
	}

	// 转换数字
	var fnums []float64
	for _, num := range nums {
		fnums = append(fnums, float64(num))
	}
	return helper(fnums)
}

func getNums(a, b float64) []float64 {
	return []float64{a + b, a - b, b - a, a * b, a / b, b / a}
}
