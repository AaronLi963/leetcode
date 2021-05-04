package main

import "sort"

func specialArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	sort.Ints(nums)
	for i := 1; i <= l; i++ {
		if nums[l-i] >= i {
			if l-i == 0 {
				return i
			} else if nums[l-i-1] < i {
				return i
			}
		}
	}
	return -1
}
