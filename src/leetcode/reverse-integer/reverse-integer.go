package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	s := ""
	radio := 1
	if x < 0 {
		radio = -1
	}
	for _, value := range strconv.Itoa(x * radio) {
		fmt.Println(value)
		s = string(value) + s
	}

	res, err := strconv.ParseInt(s,10,32)
	reverseRes := int(res) * radio
	if err == nil && reverseRes < 2147483648 && reverseRes >= -2147483648 {
		return reverseRes
	}
	return 0
}

func main() {
	fmt.Println(reverse(-120))
}
