package main

func isHappy(n int) bool {
	hash := map[int]bool{}

	for {
		sum := 0
		for n > 0 {
			mod := n % 10
			sum += mod * mod
			n /= 10
		}
		if sum == 1 {
			return true
		}
		if hash[sum] {
			break
		} else {
			hash[sum] = true
		}
		n = sum
	}

	return false
}
