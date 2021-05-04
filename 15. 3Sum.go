package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			switch {
			case sum == 0:
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
			case sum > 0:
				k--
			case sum < 0:
				j++
			}
		}
	}
	return ans
}

func main() {
	testing := [][]int{
		[]int{-1, 0, 1, 2, -1, -4},
		[]int{},
		[]int{0},
	}

	for i := range testing {
		fmt.Println(threeSum(testing[i]))
	}
}
