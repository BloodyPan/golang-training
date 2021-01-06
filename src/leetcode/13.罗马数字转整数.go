package main

import "fmt"

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func getVal(ch byte) int {
	switch ch {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func romanToInt(s string) int {
	result := 0
	preNum := getVal(s[0])
	for i := 1; i < len(s); i++ {
		num := getVal(s[i])
		if num <= preNum {
			result += preNum
		} else {
			result -= preNum
		}
		preNum = num
	}
	result += preNum

	return result
}

// @lc code=end

func main() {
	fmt.Println(romanToInt("III"))
}
