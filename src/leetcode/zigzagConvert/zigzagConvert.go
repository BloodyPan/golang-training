package main

import (
	"bytes"
	"fmt"
	"math"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	strLen := len(s)
	minUtil := 2 * numRows - 2
	radio := int(math.Floor(float64(strLen) / float64(minUtil)))
	carry := strLen % minUtil
	if carry > 0 && carry <= numRows {
		carry = 1
	} else if carry > numRows {
		carry -= numRows
	}
	
	subArrayLen := radio*(numRows-1) + carry
	resArray := make([][]rune, numRows)

	for i := 0; i < numRows; i++ {
		resArray[i] = make([]rune, int(subArrayLen))
	}
	//fmt.Println(numRows, subArrayLen, minUtil, radio, carry)

	x, y, incrX, incrY := 0, 0, 0, 0
	for _, char := range s {
		//fmt.Println(x, y)
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

type zigzag struct {
	s string
	numRows int
}

func (zz *zigzag) rowCount() int {
	if len(zz.s) < zz.numRows {
		return len(zz.s)
	}
	return zz.numRows
}

func (zz *zigzag) convert() string {
	if zz.numRows == 1 {
		return zz.s
	}

	curRow := 0
	goingDown := false

	minRows := zz.rowCount()
	resList := make([]bytes.Buffer, minRows)

	for i:=0; i<len(zz.s); i++ {
		resList[curRow].WriteByte(zz.s[i])

		if curRow == 0 || curRow == minRows - 1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow += 1
		} else {
			curRow -= 1
		}
	}
	
	s := new(bytes.Buffer)
	for _, value := range resList {
		value.WriteTo(s)
	}
	return s.String()
}

func main() {
	//fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("PAYPALISHIRING", 3))

	zz := zigzag{"PAYPALISHIRING", 4}
	fmt.Printf(zz.convert())
}
