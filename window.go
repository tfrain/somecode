package leetcode

func largestRectangleArea(heights []int) int {
	res := 0
	stack := []int{}
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i-stack[len(stack)-1]-1
			}
			area := heights[j] * width
			if area > res {
				res = area
			}
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		j := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		width := len(heights)
		if len(stack) > 0 {
			width = len(heights) - stack[len(stack)-1] -1
		}
		area := heights[j] *width
		if area > res {
			res = area
		}
	}
	return res
}

func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range s1 {
		need[s1[i]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s2) {
		r := s2[right]
		right++
		if _, ok := need[r]; ok {
			window[r]++
			if window[r] == need[r] {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(s1) {
				return true
			}
			l := s2[left]
			left++
			if _, ok := need[l]; ok {
				if window[l] == need[l] {
					valid--
				}
				window[l]--
			}
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	res := 0
	left, right := 0, 0 
	for right < len(s) {
		r := s[right]
		window[r]++
		right++
		for window[r] > 1 {
			l := s[left]
			window[l]--
			left++
		}
		if res < right-left {
			res = right-left
		}
	}
	return res
}

func minWindow(s string, t string) string {
	window, need := make(map[byte]int), make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	left, right, valid := 0, 0, 0
	leftR, rightR := 0, len(s)+1
	for right < len(s) {
		r := s[right]
		right++
		if _, ok := need[r]; ok {
			window[r]++
			if need[r] == window[r] {
				valid++
			}
		}
		for len(need) == valid {
			if right-left < rightR-leftR {
				rightR = right
				leftR = left
			}
			l := s[left]
			left++
			if _, ok := need[l]; ok {
				if need[l] == window[l] {
					valid--
				}
				window[l]--
			}
		}
	}
	if leftR == 0 && rightR == len(s)+1 {
		return ""
	}
	return s[leftR:rightR]
}

func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range p {
		need[p[i]]++
	}
	left, right, valid := 0, 0, 0
	ret := []int{}
	for right < len(s) {
		r := s[right]
		right++
		if _, ok := need[r]; ok {
			window[r]++
			if window[r] == need[r] {
				valid++
			}
		}
		if right-left >= len(p) {
			if valid == len(p) {
				ret = append(ret, left)
			}
			l := s[left]
			left++
			if _, ok := need[l]; ok {
				if window[l] == need[l] {
					valid--
				}
				window[l]--
			}
		}
	}
	return ret
}