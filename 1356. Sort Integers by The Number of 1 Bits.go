package main

import "sort"

func sortByBits(arr []int) []int {
	ones := make([][]int, 14) // 0~13
	for i := range arr {
		num := getOnes(arr[i])
		ones[num] = append(ones[num], arr[i])
	}
	ans := []int{}
	for i := range ones {
		sort.Ints(ones[i])
		ans = append(ans, ones[i]...)
	}
	return ans
}

func getOnes(i int) int {
	if i == 0 {
		return 0
	}
	count := 0
	for i > 0 {
		count += i % 2
		i /= 2
	}
	return count
}
