package leetcode

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	deque := make([][2]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				deque = append(deque, [2]int{i, j})
			}
		}
	}
	directions := [][2]int{{0,1}, {0, -1}, {-1, 0}, {1, 0}}
	minutes := 0
	for len(deque) > 0 {
		for _, d := range deque {
			deque = deque[1:]
			for _, loc := range directions {
				r, c := d[0]+loc[0], d[1]+loc[1]
				if r < 0 || r >= m || c < 0|| c >= n {
					continue
				}
				if grid[r][c] != 1 {
					continue
				}
				grid[r][c] = 2
				deque = append(deque, [2]int{r, c})
			}
		}
		if len(deque) > 0 {
			minutes++
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return minutes
}
