package main

func largestAltitude(gain []int) int {
	max := 0
	preSum := 0
	for i := range gain {
		preSum += gain[i]
		if preSum > max {
			max = preSum
		}
	}
	return max
}
