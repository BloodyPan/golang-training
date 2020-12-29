package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0  || (x % 10 == 0 && x > 0) {
		return false
	} else if x < 10 {
		return true
	}

	revert := 0
	for x > revert {
		revert = revert * 10 + x % 10
		x = x / 10
	}
	return x == revert || x == revert / 10
}

func main() {
	fmt.Println(isPalindrome(10))
}
