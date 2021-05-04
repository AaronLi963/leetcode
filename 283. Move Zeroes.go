package main

func moveZeroes(nums []int) {
	var i, p int
	for p < len(nums) {
		if nums[p] == 0 {
			p++
		} else {
			temp := nums[p]
			nums[p] = nums[i]
			nums[i] = temp
			i++
			p++
		}
	}
}
