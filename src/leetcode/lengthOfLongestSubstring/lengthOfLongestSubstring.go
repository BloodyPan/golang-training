package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var length int = 0
	var charIndexMap map[rune]int = make(map[rune]int)
	i := 0
	for index, char := range s {
		charIndex := charIndexMap[char]
		if charIndex > i {
			i = charIndex
		}
		currentLen := index - i + 1
		if currentLen > length {
			length = currentLen
		}
		charIndexMap[char] = index + 1
	}
	return length
}

func main() {
	str := "abcabcabc"
	fmt.Printf("%d", lengthOfLongestSubstring(str))
}
