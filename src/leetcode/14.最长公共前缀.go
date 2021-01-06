package main

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
import (
	"bytes"
)

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	preStr := strs[0]
	for i := 1; i < len(strs); i++ {
		commonPrefix := bytes.Buffer{}
		minLen := min(len(strs[i]), len(preStr))
		for j := 0; j < minLen; j++ {
			if strs[i][j] == preStr[j] {
				commonPrefix.WriteByte(preStr[j])
			} else {
				break
			}
		}
		preStr = commonPrefix.String()
	}
	return preStr
}

// @lc code=end

func main() {
	longestCommonPrefix([]string{"flower", "flow", "flight"})
	longestCommonPrefix([]string{"cir", "car"})
	//fmt.Println(longestCommonPrefix([]string{"cir", "car"}))
}
