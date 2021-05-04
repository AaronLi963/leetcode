package main

func subarraySum(nums []int, k int) int {
	l := len(nums)
	sum := 0
	m := map[int]int{}
	m[0] = 1
	ans := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
		ans += m[sum-k]
		m[sum]++
	}
	return ans

}
