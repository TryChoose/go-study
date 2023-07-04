package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var dp [105][105]int
	var aa [105][105]int
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Scan(&aa[i][j])
		}
	}
	for i := 1; i <= n; i++ {
		dp[n][i] = aa[n][i]
	}
	for i := n - 1; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			dp[i][j] = max(dp[i+1][j], dp[i][j+1]) + aa[i][j]
		}
	}
	fmt.Println(dp[1][1])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
