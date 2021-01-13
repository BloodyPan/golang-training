package main

import (
	"container/ring"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
	ring.Ring{}
}

// @lc code=end

func main() {
	longestCommonPrefix([]string{"flower", "flow", "flight"})
	longestCommonPrefix([]string{"cir", "car"})
	fmt.Println(longestCommonPrefix([]string{"cir", "car"}))
	fmt.Println(longestCommonPrefix([]string{"cir", "car"}))

	a, b := 5, 6
	fmt.Printf("%d %d\n", a, b)

	a, b = b, a
	fmt.Printf("%d %d", a, b)
}
