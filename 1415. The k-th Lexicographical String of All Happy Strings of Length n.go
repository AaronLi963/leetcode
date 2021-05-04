package main

import "math"

func getHappyString(n int, k int) string {
	totalCase := int(3 * math.Pow(float64(2), float64(n-1)))
	if totalCase < k {
		return ""
	}

	div := make([]int, n)
	div[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		div[i] = div[i+1] * 2
	}
	k--
	char := []byte{'a', 'b', 'c'}
	ans := make([]byte, n)

	for i := 0; i < n; i++ {
		ans[i] = char[k/div[i]]
		switch char[k/div[i]] {
		case 'a':
			char = []byte{'b', 'c'}
		case 'b':
			char = []byte{'a', 'c'}
		case 'c':
			char = []byte{'a', 'b'}
		}
		k %= div[i]
	}

	return string(ans)
}
