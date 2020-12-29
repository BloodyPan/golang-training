package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	sign := 1
	res := new(bytes.Buffer)
	for i, v := range s {
		if v >= '0' && v <= '9' {
			res.WriteString(string(v))
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}
	}
	result, err := strconv.Atoi(res.String())
	result = result * sign
	if result < -2147483648 {
		return -2147483648
	}
	if result > 2147483647 {
		return 2147483647
	}
	if err == nil {
		return result
	}
	return 0
}

func main() {
	fmt.Println(myAtoi("aaa -5"))
	fmt.Println(strings.TrimSpace(" asdas 1312 "))
}

