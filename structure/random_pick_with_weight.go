package structure

import "math/rand"

type randSolution struct {
	preSum []int
}

func Constructor(w []int) randSolution {
	n := len(w)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + w[i-1]
	}
	return randSolution{
		preSum: preSum,
	}
}

func (this *randSolution) PickIndex() int {
	n := len(this.preSum)
	target := rand.Intn(this.preSum[n-1]) + 1
	l, r := 0, n
	for l < r {
		m := l + (r-l)/2
		if this.preSum[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l - 1
}
