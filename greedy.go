package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	n, sum, minSum, start := len(gas), 0, 0, 0
	for i := range gas {
		sum += gas[i]-cost[i]
		if sum < minSum {
			minSum = sum
			start = i+1
		}
	}
	if sum < 0 {
		return -1
	}
	if start == n {
		return 0
	}
	return start
}

