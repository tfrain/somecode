package leetcode

import "math"

// 代表性, 需要遍历dp数组(不遍历显然可以)
func longestValidParentheses(s string) int {
	dp := make([]int, len(s)+1)
	stack := make([]int, 0)
	res := 0
	for i := range s {
		if s[i] == '[' {
			dp[i+1] = 0
			stack = append(stack, i)
		} else {
			if len(stack) > 0 {
				j := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				dp[i+1] = i - j + 1 + dp[j]
				if dp[i+1] > res {
					res = dp[i+1]
				}
			} else {
				dp[i+1] = 0
			}
		}
	}
	return res
}

// 背包问题
// 背包问题比简单的dp要复杂, 如爬楼梯2, 中间的状态从后面得到, 如果不行就一定是false, 所以可以简单dp table 或备忘录
// [1,2,3,5,17,6,14,12,6] 一半是33, 不可行
// 32、30、27、22、5; ---》dp5 不行---》22(dp16、2、8、10、4)(6,14,12,6) 不行, 27(dp 21、7、1、13、15、3、9)(17,6,14,12,6) 不行,
// 30(dp 25、19、24、10、12、6、16)(5,17,6,14,12,6) 不行
// 从31(33-2)(1,3,5,17,6,14,12,6)--》28(dp 23, 6, 0)(5,17,6) 这里不行,dp6为0了, 使用start 不行
// 和第 322题 coin change 进行比较, 那里可以用普通递归, 是因为每个数字可重复, 某memo[v]不能过, 确实不能通过, 和爬楼梯类似;
// 背包问题中元素是否可重复, 会导致完全不同的结果
func canPartitionFail(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 0 {
		sum = sum / 2
	} else {
		return false
	}
	memo := make([]int, sum+1)
	var dp func(amount, start int) bool
	dp = func(amount, start int) bool {
		if amount == 0 {
			return true
		}
		if amount < 0 || start >= len(nums) {
			return false
		}
		if memo[amount] != 0 {
			return memo[amount] == 1
		}
		for i := start; i < len(nums); i++ {
			if amount-nums[i] < 0 {
				continue
			}
			if dp(amount-nums[i], i+1) {
				memo[amount] = 1
				return true
			}
		}
		memo[amount] = -1
		return false
	}
	return dp(sum, 0)
}

// [1,2,3,5,17,6,14,12,6] 一半是33, 不可行,
// 32、30、27、22、5; ---》由上可知dp(16、2、8、10、4、21、7、1、13、15、3、9、25、19、24、10、12、6、16) 都是false
// 所以32 应该也不行; 从31(33-2)(1,3,5,17,6,14,12,6)--》28(dp 23,6)(5,17,6) 不行,dp6 为false, 都不可行
// 16(33-17)(dp 14 2)(2,14)
// 此例子使用used可行

// 但[1,4,5,8,12,13,14,17,18] 一半的数字是 46, 不可行;
// [1,13,14,18] 可行, 肯定是中间又被覆盖,所以不可行
// 使用used不可行
func canPartitionFail1(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 0 {
		sum = sum / 2
	} else {
		return false
	}
	memo := make([]int, sum+1)
	used := make([]bool, len(nums))
	var dp func(amount int) bool
	dp = func(amount int) bool {
		if amount == 0 {
			return true
		}
		if amount < 0 {
			return false
		}
		if memo[amount] != 0 {
			return memo[amount] == 1
		}
		for i := range nums {
			if amount-nums[i] < 0 {
				continue
			}
			used[i] = true
			if dp(amount - nums[i]) {
				memo[amount] = 1
				return true
			}
			used[i] = false
		}
		memo[amount] = -1
		return false
	}
	return dp(sum)
}

func canPartition(nums []int) bool {
	sum, n := 0, len(nums)
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	} else {
		sum = sum / 2
	}
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, sum+1)
	}
	var dp func(amount, start int) bool
	dp = func(amount, start int) bool {
		if amount == 0 {
			return true
		}
		if start < 0 || start >= n || amount < 0 {
			return false
		}
		if memo[start][amount] != 0 {
			return memo[start][amount] == 1
		}
		if dp(amount, start+1) || dp(amount-nums[start], start+1) {
			memo[start][amount] = 1
			return true
		} else {
			memo[start][amount] = -1
			return false
		}

	}
	return dp(sum, 0)
}

func canPartition1(nums []int) bool {
	sum, n := 0, len(nums)
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	} else {
		sum = sum / 2
	}
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] < 0 {
				// 背包容量不足，不能装入第 i 个物品; 忽略它, 维系状态
				// 这个物品不适合背包, 但是其他物品可能合适, false
				// 这个物品不需要装, 之前的物品已经使其满了, true
				dp[i][j] = dp[i-1][j]
			} else {
				// 能装下, 从之前的状态判断
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[n][sum]
}

// 代表性, 返回dp(0)
func jump(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = n
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var dp func(start int) int
	dp = func(start int) int {
		if start >= n-1 {
			return 0
		}
		if memo[start] != n {
			return memo[start]
		}
		res := math.MaxInt32
		for i := 1; i <= nums[start]; i++ {
			res = min(res, dp(i)+1)
		}
		if res == math.MaxInt32 {
			memo[start] = -1
			return -1
		}
		memo[start] = res
		return res
	}
	return dp(0)
}

func jump1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = n - 1
	}
	dp[n-1] = 0
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j <= nums[i] && j < n; j++ {
			if dp[i] > dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[0]
}

func canJump(nums []int) bool {
	n := len(nums)
	memo := make([]int, n)
	var dp func(start int) bool
	dp = func(start int) bool {
		if start >= n-1 {
			return true
		}
		if memo[start] != 0 {
			return memo[start] == 1
		}
		for i := 1; i <= nums[start]; i++ {
			if dp(start + i) {
				memo[start] = 1
				return true
			}
		}
		memo[start] = -1
		return false
	}
	return dp(0)
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	ret := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if ret < dp[i] {
			ret = dp[i]
		}
	}
	return ret
}

// 复杂背包问题
func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var dp func(amount int) int
	dp = func(amount int) int {
		if amount == 0 {
			return 0
		}
		if memo[amount] != 0 {
			return memo[amount]
		}
		res := math.MaxInt32
		for _, coin := range coins {
			if amount-coin < 0 {
				continue
			}
			subNum := dp(amount-coin) + 1
			res = min(res, subNum)
		}
		if res == math.MaxInt32 {
			memo[amount] = -1
			return -1
		}
		memo[amount] = res
		return memo[amount]
	}
	return dp(amount)
}

func coinChange0(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 0
	}
	// 除了第一列, 都变成amount+1, 方便计算
	for i := 0; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = amount+1
		}
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				// 可重复, 所以也可能用dp[i], 比如1, 5, 5
				dp[i][j] = min(min(dp[i-1][j], dp[i-1][j-coins[i-1]]+1), dp[i][j-coins[i-1]]+1)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	if dp[n][amount] == amount+1 {
		return -1
	}
	return dp[n][amount]
}

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	dp[0] = 0
	for i := range dp {
		for _, v := range coins {
			if i == 0 || i-v < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-v]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func change(amount int, coins []int) int {
    n := len(coins)
    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, amount+1)
        dp[i][0] = 1
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= amount; j++ {
            if j - coins[i-1] >= 0 {
                dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
    return dp[n][amount]
}

func rob(nums []int) int {
	n := len(nums)
	if n < 2 {
		return nums[0]
	}
	dp := make([]int, n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp[0] = nums[0]
	dp[1] = max(nums[1], dp[0])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}
	min := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// base case, 一方为0,则直接插入另一方的长度
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i][j-1] > dp[i-1][j] {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[m][n]
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// 备忘录和dp table 的方法掌握还是不够熟练
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	memo := make([]int, n)
	var dp func(start int) bool
	dp = func(start int) bool {
		if start > n-1 {
			return true
		}
		if memo[start] != 0 {
			return memo[start] == 1
		}
		for _, word := range wordDict {
			wordLen := len(word)
			if start+wordLen > n {
				continue
			}
			if word != s[start:start+wordLen] {
				continue
			}
			if dp(start + wordLen) {
				memo[start] = 1
				return true
			}
		}
		memo[start] = -1
		return false
	}
	return dp(0)
}
