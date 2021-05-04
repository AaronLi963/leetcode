package main

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	ans := [][]int{}
	minDiff := 20000001
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff == minDiff {
			ans = append(ans, []int{arr[i], arr[i+1]})
		} else if diff < minDiff {
			minDiff = diff
			ans = [][]int{[]int{arr[i], arr[i+1]}}
		}
	}
	return ans
}
