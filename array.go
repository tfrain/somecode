package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	zeroIdx := make([][2]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroIdx = append(zeroIdx, [2]int{i, j})
			}
		}
	}
	if len(zeroIdx) == 0 {
		return
	}
	setZero := func(row, col int) {
		for i := 0; i < m; i++ {
			matrix[i][col] = 0
		}
		for i := 0; i < n; i++ {
			matrix[row][i] = 0
		}
	}
	for _, idx := range zeroIdx {
		setZero(idx[0], idx[1])
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	startGraph := make([][]int, numCourses)
	peakDegree := make([]int, numCourses)
	for _, p := range prerequisites {
		startGraph[p[1]] = append(startGraph[p[1]], p[0])
		peakDegree[p[0]]++
	}
	notPointedList := make([]int, 0)
	for _, v := range peakDegree {
		if v == 0 {
			notPointedList = append(notPointedList, v)
		}
	}
	for i := 0; i < len(notPointedList); i++ {
		for _, p := range startGraph[notPointedList[i]] {
			peakDegree[p]--
			if peakDegree[p] == 0 {
				notPointedList = append(notPointedList, p)
			}
		}
	}
	return len(notPointedList) == numCourses
}

func twoSum(nums []int, target int) []int {
	hasMap := make(map[int]int, 0)
	for idx, num := range nums {
		if v, ok := hasMap[target-num]; ok {
			return []int{idx, v}
		}
		hasMap[num] = idx
	}
	return []int{}
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	var bt func(start int, sub []int)
	bt = func(start int, sub []int) {
		if start > n {
			return
		}
		tmp := make([]int, len(sub))
		copy(tmp, sub)
		res = append(res, tmp)
		res = append(res, sub)
		for i := start + 1; i < n; i++ {
			bt(i, append(sub, nums[i]))
		}
	}
	bt(0, nil)
	return res
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	prefix := 1
	for i, v := range nums {
		res[i] = prefix
		prefix *= v
	}
	prefix = 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= prefix
		prefix *= nums[i]
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	strsMap := make(map[string][]string)
	for _, str := range strs {
		code := encode(str)
		strsMap[code] = append(strsMap[code], str)
	}
	ret := [][]string{}
	for _, vals := range strsMap {
		ret = append(ret, vals)
	}
	return ret
}

func encode(str string) string {
	vals := make([]byte, 26)
	for i := range str {
		c := str[i] - 'a'
		vals[c]++
	}
	return string(vals)
}

func decodeString(s string) string {
	n := len(s)
	var decode func(start int) (string, int)
	decode = func(start int) (string, int) {
		if start >= n {
			return "", 0
		}
		num := 0
		var res string
		for i := start; i < n; i++ {
			if n, err := strconv.Atoi(string(s[i])); err == nil {
				num = num*10 + n
			} else if s[i] == '[' {
				subs, end := decode(i + 1)
				res += strings.Repeat(subs, num)
				num = 0
				i = end + 1
			} else if s[i] == ']' {
				return res, i
			} else {
				res += string(s[i])
			}
		}
		return res, 0
	}
	ret, _ := decode(0)
	return ret
}

func subarraySum(nums []int, k int) int {
	prefix := map[int]int{0: 1}
	sum, cnt := 0, 0
	for _, v := range nums {
		sum += v
		cnt += prefix[sum-k]
		prefix[sum]++
	}
	return cnt
}

func partitionLabels(s string) []int {
	var ret []int
	need := make(map[byte]int)
	for i := range s {
		need[s[i]] = i
	}
	left, right := 0, 0
	max := 0
	for right < len(s) {
		if max < need[s[right]] {
			max = need[s[right]]
		}
		if right == max {
			ret = append(ret, right-left+1)
			max = 0
			left = right + 1
		}
	}
	return ret
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if l := len(nums1) + len(nums2); l%2 == 0 {
		return (findKth(nums1, nums2, l/2-1) + findKth(nums1, nums2, l/2)) / 2
	} else {
		return findKth(nums1, nums2, l/2)
	}
}

func findKth(nums1, nums2 []int, k int) float64 {
	for {
		l1, l2 := len(nums1), len(nums2)
		m1, m2 := l1/2, l2/2
		if l1 == 0 {
			return float64(nums2[k])
		} else if l2 == 0 {
			return float64(nums1[k])
		} else if k == 0 {
			if n1, n2 := nums1[0], nums2[0]; n1 > n2 {
				return float64(n2)
			} else {
				return float64(n1)
			}
		}
		if k <= m1+m2 {
			if nums1[m1] > nums2[m2] {
				nums1 = nums1[:m1]
			} else {
				nums2 = nums2[:m2]
			}
		} else {
			if nums1[m1] > nums2[m2] {
				nums2 = nums2[m2+1:]
				k -= m2 + 1
			} else {
				nums1 = nums1[m1+1:]
				k -= m1 + 1
			}
		}
	}
}

func search(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for l <= h {
		m := l + (h - l) + 1
		if nums[m] < target {
			if target >= nums[0] && nums[m] < nums[0] {
				h = m - 1
			} else {
				l = m + 1
			}
		} else if nums[m] > target {
			if nums[m] >= nums[0] && target < nums[0] {
				l = m + 1
			} else {
				h = m - 1
			}
		} else {
			return nums[m]
		}
	}
	return -1
}

func majorityElement(nums []int) int {
	target, cnt := nums[0], 1
	for _, v := range nums {
		if v == target {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			target = v
			cnt = 1
		}
	}
	return target
}

func sortColors(nums []int) {
	zero, one, two := -1, 0, len(nums)
	for one < two {
		if nums[one] == 0 {
			zero++
			nums[zero], nums[one] = nums[one], nums[zero]
			one++
		} else if nums[one] == 2 {
			two--
			nums[one], nums[two] = nums[two], nums[one]
		} else {
			one++
		}
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true
		}
	}
	return false
}

func merge(intervals [][]int) [][]int {
	var ret [][]int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ret = append(ret, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := ret[len(ret)-1]
		if intervals[i][0] < last[1] {
			last[1] = max(last[1], intervals[i][1])
		} else {
			ret = append(ret, intervals[i])
		}
	}
	return ret
}

func generate(numRows int) [][]int {
	ret := [][]int{{1}}
	for i := 2; i < numRows; i++ {
		lastRow := ret[len(ret)-1]
		oneRow := []int{1}
		for i := 1; i < len(lastRow); i++ {
			oneRow = append(oneRow, lastRow[i-1]+lastRow[i])
		}
		oneRow = append(oneRow, 1)
		ret = append(ret, oneRow)
	}
	return ret
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
	}
	return l
}

func maxProfit(prices []int) int {
	minPrice, res := prices[0], 0
	for _, v := range prices {
		if v > minPrice {
			res = v - minPrice
		} else {
			minPrice = v
		}
	}
	return res
}

func longestConsecutive(nums []int) int {
	count := make(map[int]struct{}, 0)
	for _, num := range nums {
		count[num] = struct{}{}
	}
	max := 0
	for num := range count {
		if _, ok := count[num-1]; ok {
			continue
		}
		currNum, currLen := num, 1
		_, ok := count[currNum+1]
		for ok {
			currNum++
			currLen++
			_, ok = count[currNum+1]
		}
		if currLen > max {
			max = currLen
		}
	}
	return max
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] >= nums[0] {
			l = m + 1
		} else {
			if nums[m] < nums[m-1] {
				return nums[m]
			} else {
				r = m - 1
			}
		}
	}
	return -1
}

// 167
func twoSum1(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l <= r {
		sum := numbers[l] + numbers[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return nil
}
