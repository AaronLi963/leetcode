package main

import "sort"

func twoCitySchedCost(costs [][]int) int {
	diff := make([]int, len(costs))
	sum := 0
	for i, cost := range costs {
		diff[i] = cost[0] - cost[1]
		sum += cost[0]
	}
	sort.Ints(diff)
	for i := len(diff) / 2; i < len(diff); i++ {
		sum -= diff[i]
	}
	return sum
}
