package leetcode

import (
	"math"
	"sort"
)

// https://docs.aws.amazon.com/codewhisperer/latest/userguide/actions-and-shortcuts.html
func longestPalindrome(s string) string {
	palindrome := func(i, j int) string {
		for s[i] == s[j] && i >= 0 && j < len(s) {
			i--
			j++
		}
		return s[i+1 : j]
	}
	var res string
	for i := 0; i < len(s)-1; i++ {
		first := palindrome(i, i)
		if len(first) > len(res) {
			res = first
		}
		second := palindrome(i, i+1)
		if len(second) > len(res) {
			res = second
		}
	}
	return res
}

func maxProduct(nums []int) int {
	maxV, minV, res := nums[0], nums[0], nums[0]
	max, min := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}, func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if v < 0 {
			maxV, minV = minV, maxV
		}
		maxV = max(maxV*v, v)
		minV = min(minV*v, v)
		res = max(res, maxV)
	}
	return res
}

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return nTargetSum(nums, 3, 0, 0)
}

func nTargetSum(nums []int, n, target, start int) [][]int {
	var ret [][]int
	size := len(nums)
	if n > size {
		return ret
	}
	if n == 2 {
		l, r := start, size-1
		for l < r {
			left, right := nums[l], nums[r]
			sum := left + right
			if sum > target {
				for l < r && nums[l] == left {
					l++
				}
			} else if sum < target {
				for l < r && nums[r] == right {
					r--
				}
			} else {
				ret = append(ret, []int{left, right})
				for l < r && nums[l] == left {
					l++
				}
				for l < r && nums[r] == right {
					r--
				}
			}
		}
	} else {
		for i := start; i < size; i++ {
			subs := nTargetSum(nums, n-1, target-nums[i], i+1)
			for _, sub := range subs {
				one := []int{nums[i]}
				one = append(one, sub...)
				ret = append(ret, one)
			}
			for i < size-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return ret
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse := func(i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	reverse(0, n-1)
	reverse(0, k-1)
	reverse(k, n-1)
}
func romanToInt1(s string) int {
	getInt := func(v byte) int {
		switch v {
		case 'I':
			return 1
		case 'V':
			return 5
		case 'X':
			return 10
		case 'L':
			return 50
		case 'C':
			return 100
		case 'D':
			return 500
		case 'M':
			return 1000
		default:
			return -1
		}
	}
	ret, prev := 0, 0
	for i := range s {
		now := getInt(s[i])
		ret += now
		if now > prev {
			ret -= 2 * prev
		}
		prev = now
	}
	return ret
}

func romanToInt(s string) int {
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	stack := make([]int, 0)
	res := 0
	for i := range s {
		if len(stack) != 0 && m[s[len(stack)-1]] > m[s[i]] {
			// jisuan
			res += compute(m, s[stack[0]:stack[len(stack)]])
			stack = stack[:0]
		}
		stack = append(stack, i)
	}
	if len(stack) != 0 {
		res += compute(m, s[stack[0]:stack[len(stack)]])
	}
	return res
}

func compute(m map[byte]int, s string) int {
	res := m[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		res -= m[s[i]]
	}
	return res
}

// ["flower","flower","flower","flower"]
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	minLen := math.MaxInt32
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	idx, n := 0, len(strs)
	ret := []byte{}
	for idx < minLen {
		c := strs[0][idx]
		for i := 1; i < n; i++ {
			if c != strs[i][idx] {
				return string(ret)
			}
		}
		ret = append(ret, c)
		idx++
	}

	return string(ret)
}

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(prefix) {
			prefix = prefix[:len(strs[i])]
		}
		var j int
		for ; j < len(prefix); j++ {
			if strs[i][j] != prefix[j] {
				break
			}
		}
		prefix = prefix[:j]
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	var res []byte
	idx := 0
	for idx < len(strs[0]) {
		flag := true
		c := strs[0][idx]
		for _, str := range strs[1:] {
			if len(str) <= idx || str[idx] != c {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		res = append(res, c)
		idx++
	}
	return string(res)
}

func addStrings(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	var prev byte
	var res []byte
	for i, j := n1-1, n2-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		oneRoundSum := prev
		if i >= 0 {
			oneRoundSum += num1[i] - '0'
		}
		if j >= 0 {
			oneRoundSum += num2[j] - '0'
		}
		prev = oneRoundSum / 10
		res = append(res, '0'+oneRoundSum%10)
	}
	if prev > 0 {
		res = append(res, '1')
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 1 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return -1
}

func isValid(s string) bool {
	// match := func(c byte) byte {
	// 	if c == ']' {
	// 		return  '['
	// 	}
	// 	if c == ')' {
	// 		return '('
	// 	}
	// 	return '{'
	// }
	stack := []byte{}
	return len(stack) == 0
}

func firstUniqChar(s string) int {
	var cnt [26]byte
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if cnt[s[i]-'a'] == '1' {
			return i
		}
	}
	return -1
}

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	res := 1
	for i := 2; i < num/i; i++ {
		if num%i == 0 {
			res += i + num/i
		} else {
			continue
		}
	}
	return res == num
}

func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && (num-1)%3 == 0
}

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if target < nums[m] {
			r = m - 1
		} else if target > nums[m] {
			l = m + 1
		} else {
			minIdx, maxIdx := m, m
			for i := 0; i <= m; i++ {
				if nums[i] == target {
					minIdx = i
					i = m
				}
			}
			for j := len(nums) - 1; j >= m; j-- {
				if nums[j] == target {
					maxIdx = j
					j = m
				}
			}
			return []int{minIdx, maxIdx}
		}
	}
	return []int{-1, -1}
}

func countAndSay(n int) string {
	var recursion func(b []byte) []byte
	recursion = func(b []byte) []byte {
		var buf []byte
		for i, j := 0, 1; i < len(b); i = j {
			// 11 -> 21
			for j < len(b) && b[i] == b[j] {
				j++
			}
			// 将差值带入下一轮
			buf = append(buf, uint8(j-i+'0'), b[i])
		}
		return buf
	}
	b := []byte{'1'}
	for n > 1 {
		b = recursion(b)
		n--
	}
	return string(b)
}

func trap(height []int) int {
	maxL, maxR, res := 0, 0, 0
	n := len(height)
	left, right := 0, n-1
	max := func(a, b int) int{
		if a > b {
			return a
		}
		return b
	}
	for left <= right {
		maxL = max(maxL, height[left])
		maxR = max(maxR, height[right])
		if maxL < maxR {
			res += maxL-height[left]
			left++
		} else {
			res += maxR-height[right]
			right--
		}
	}
	return res
}