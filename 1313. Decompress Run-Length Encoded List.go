package main

func decompressRLElist(nums []int) []int {
	ans := []int{}
	for i := 0; i < len(nums)/2; i++ { // 2n, 2n+1
		for j := 0; j < nums[2*i]; j++ {
			ans = append(ans, nums[2*i+1])
		}
	}
	return ans
}
