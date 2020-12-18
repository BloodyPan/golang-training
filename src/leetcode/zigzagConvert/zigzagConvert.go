package main

import (
	"fmt"
	"math"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	strLen := len(s)
	minUtil := 2 * numRows - 2
	radio := int(math.Ceil(float64(strLen) / float64(minUtil)))
	carry := strLen % minUtil
	if carry > 0 && carry <= 4 {
		carry = 1
	} else if carry > 4 {
		carry -= 4
	}

	subArrayLen := radio*(numRows-1) + carry
	resArray := make([][]rune, numRows)

	for i := 0; i < numRows; i++ {
		resArray[i] = make([]rune, int(subArrayLen))
	}
	fmt.Println(numRows, subArrayLen, minUtil, radio, carry)

	x, y, incrX, incrY := 0, 0, 0, 0
	for _, char := range s {
		fmt.Println(x, y)
		resArray[x][y] = char


		if x == numRows - 1 {
			incrX = -1
			incrY = 1
		} else if x == 0 {
			incrX = 1
			incrY = 0
		}

		x += incrX
		y += incrY
		if y >= subArrayLen {
			break
		}
	}

	res := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < subArrayLen; j++ {
			if resArray[i][j] != 0 {
				res += string(resArray[i][j])
			}
		}
	}
	fmt.Println(resArray)
	return res
}

func main() {
	//fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("ABCDE", 4))
}
