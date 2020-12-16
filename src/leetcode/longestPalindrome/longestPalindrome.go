package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for l := 0; l < n; l++ {
		for i:= 0; i < n; i ++ {
			j := i+l
			if j >= n {
				break
			}
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			if dp[i][j] && l + 1 > len(ans) {
				ans = s[i: j+1]
			}
		}
	}

	return ans
}

func main() {
	s := "asdadsss"
	fmt.Printf("%s: %s", s, longestPalindrome(s))
}
