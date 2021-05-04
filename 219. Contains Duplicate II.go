package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	if len(nums) <= 1 {
		return false
	}
	if k >= len(nums) {
		k = len(nums) - 1
	}
	var i, j int
	for i+j+1 <= k+1 {
		m[nums[j]]++
		if m[nums[j]] > 1 {
			return true
		}
		j++
	}
	j--
	for j < len(nums)-1 {
		m[nums[i]]--
		i++
		j++
		m[nums[j]]++
		if m[nums[j]] > 1 {
			return true
		}
	}
	return false
}
