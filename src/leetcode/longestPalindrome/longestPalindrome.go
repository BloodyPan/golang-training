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

// region 中心扩展
func longestPalindromeCenter(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i:=0; i<len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1 - left1 > end - start {
			start, end = left1, right1
		}
		if right2 - left2 > end - start {
			start, end = left2, right2
		}
	}
	return s[start: end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ;left>=0 && right < len(s) && s[left] == s[right];left, right = left - 1, right + 1 {}
	return left + 1, right - 1
}
// endregion

func main() {
	s := "asdadsss"
	fmt.Printf("%s: %s\n", s, longestPalindrome(s))
	fmt.Printf("%s: %s", s, longestPalindromeCenter(s))
}
