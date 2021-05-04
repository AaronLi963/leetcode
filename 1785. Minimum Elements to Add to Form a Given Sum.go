package main

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if goal == sum {
		return 0
	}
	return (abs(goal-sum)-1)/limit + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
