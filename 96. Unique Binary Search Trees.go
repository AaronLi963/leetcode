package main

func numTrees(n int) int {
	dp := []int{1, 1}
	if n < 2 {
		return dp[n]
	}
	for i := 2; i <= n; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += dp[j] * dp[i-j-1]
		}
		dp = append(dp, sum)
	}
	return dp[n]
}
