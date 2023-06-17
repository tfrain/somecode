package leetcode

func permute(nums []int) [][]int {
	n := len(nums)
	used := make([]bool, n)
	ret := make([][]int, 0)
	var bt func(subs []int)
	bt = func(subs []int) {
		if len(subs) == n {
			ret = append(ret, subs)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			bt(append(subs, nums[i]))
			used[i] = false
		}
	}
	bt(nil)
	return ret
}

func partition(s string) [][]string {
	ret := [][]string{}
	for i := range s {
		if sub := s[:i+1]; check(sub, 0, len(sub)-1) {
			subs := partition(s[i+1:])
			for _, v := range subs {
				ret = append(ret, append([]string{sub}, v...))
			}
		}
	}
	return ret
}

func partition1(s string) [][]string {
	var ret [][]string
	var bt func(subs []string, str string)
	bt = func(subs []string, str string) {
		if len(str) == 0 {
			tmp := make([]string, len(subs))
			copy(tmp, subs)
			ret = append(ret, subs)
			return
		}
		for i := 0; i < len(str); i++ {
			if substr := str[:i+1]; check(substr, 0, len(substr)-1) {
				bt(append(subs, substr), str[i+1:])
			}
		}
	}
	bt(nil, s)
	return ret
}

func check(s string, i, j int) bool {
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func solveNQueens(n int) [][]string {
	fields := make([][]byte, n)
	for i := range fields {
		fields[i] = make([]byte, n)
	}
	for i := range fields {
		for j := range fields[i] {
			fields[i][j] = '.'
		}
	}
	var ret [][]string
	var bt func(int)
	bt = func(row int) {
		if row == n {
			ret = append(ret, toString(fields))
			return
		}
		for i := 0; i < n; i++ {
			if !validN(fields, row, i) {
				continue
			}
			fields[row][i] = 'Q'
			bt(row + 1)
			fields[row][i] = '.'
		}
	}
	bt(0)
	return ret
}

func validN(fields [][]byte, row, col int) bool {
	n := len(fields)
	for i := 0; i < n; i++ {
		if fields[row][i] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if fields[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if fields[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func toString(fields [][]byte) (res []string) {
	for i := range fields {
		res = append(res, string(fields[i]))
	}
	return
}

func letterCombinations(digits string) []string {
	var ret []string
	letter := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var bt func([]byte)
	bt = func(subs []byte) {
		if len(subs) == len(digits) {
			ret = append(ret, string(subs))
			return
		}
		idx := digits[len(subs)] - '2'
		for i := 0; i < len(letter[idx]); i++ {
			bt(append(subs, letter[idx][i]))
		}
	}
	bt(nil)
	return ret
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var combination func(subs []int, sum, start int)
	combination = func(subs []int, sum, start int) {
		if sum == 0 {
			tmp := make([]int, len(subs))
			copy(tmp, subs)
			res = append(res, tmp)
			return
		}
		if sum < 0 {
			return
		}
		for i := start; i < len(candidates); i++ {
			if sum-candidates[i] < 0 {
				continue
			}
			combination(append(subs, candidates[i]), sum-candidates[i], i)
		}
	}
	combination(nil, target, 0)
	return res
}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var bfs func(i, j int)
	bfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		for _, d := range directions {
			l, r := d[0]+i, d[1]+j
			bfs(l, r)
		}
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
				bfs(i, j)
			}
		}
	}
	return res
}

func generateParenthesis(n int) []string {
	var res []string
	var bt func(subs []byte, left, right int)
	bt = func(subs []byte, left, right int) {
		if left == 0 && right == 0 {
			res = append(res, string(subs))
			return
		}
		if left > right || left < 0 || right < 0 {
			return
		}
		bt(append(subs, '('), left-1, right)
		bt(append(subs, ')'), left, right-1)
	}
	bt(nil, n, n)
	return res
}
