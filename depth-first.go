package leetcode

// 更像是回溯, 进行剪枝, 而且也更像是广度优先, 而不是深度优先
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	directions := [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	var bt func(start, i, j int) bool
	bt = func(start, i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if board[i][j] != word[start] || used[i][j] {
			return false
		}
		for _, d := range directions {
			r, c := i+d[0], j+d[1]
			used[r][c] = true
			if bt(start+1, r, c) {
				return true
			}
			used[r][c] = false
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if bt(0, i, j) {
				return true
			}
		}
	}
	return false
}
