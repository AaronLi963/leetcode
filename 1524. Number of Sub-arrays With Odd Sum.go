package main

func numOfSubarrays(arr []int) int {
	ans := 0
	var odds, evens int

	for i := range arr {
		if arr[i]%2 == 0 {
			evens++
		} else {
			swap(&odds, &evens)
			odds++
		}

		ans += odds
	}

	return ans % 1000000007
}
