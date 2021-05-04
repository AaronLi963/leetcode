package main

func checkPowersOfThree(n int) bool {
	pows := make([]int, 15)
	pows[0] = 1
	for i := 1; i < 15; i++ {
		pows[i] = pows[i-1] * 3
	}
	for i := 14; i >= 0; i-- {
		if n >= pows[i] {
			n -= pows[i]
		}
		if n == 0 {
			return true
		}
	}
	return false
}
