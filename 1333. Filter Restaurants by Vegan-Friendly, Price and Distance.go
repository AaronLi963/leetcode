package main

import "sort"

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	rates := []int{}
	ratingMap := map[int][]int{} // rating : ids

	for i := range restaurants {
		if restaurants[i][2] >= veganFriendly {
			if restaurants[i][3] <= maxPrice && restaurants[i][4] <= maxDistance {
				ratingMap[restaurants[i][1]] = append(ratingMap[restaurants[i][1]], restaurants[i][0])
				rates = append(rates, restaurants[i][1])
			}
		}
	}
	sort.Ints(rates)
	ans := []int{}
	for i := len(rates) - 1; i >= 0; i-- {
		if i < len(rates)-1 {
			if rates[i] == rates[i+1] {
				continue
			}
		}
		ids := ratingMap[rates[i]]
		sort.Ints(ids)
		for j := len(ids) - 1; j >= 0; j-- {
			ans = append(ans, ids[j])
		}

	}
	return ans
}
