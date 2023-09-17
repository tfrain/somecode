package interview

// case1: abc ahbgdc
// case2: axc ahbgdc
// 暴力法，flag变量； valToIdx 可以做, 不要想到dp上去了
func substring(s, t string) bool {
	idx := -1
	for i := range s {
		flag := false
		for j := range t {
			if s[i] == t[j] { // 这里写成s=s了
				if j <= idx {
					return false
				}
				idx = j
				flag = true
				// break 不能用这个
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
