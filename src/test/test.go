package main

import (
	"fmt"
)

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		defer func() {
			a, b = b, a+b
		}()
		return a
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go func() { fmt.Println(i) }()
	}

	// for i := 0; i < 100; i++ {
	// 	j := i
	// 	go func() {
	// 		fmt.Println(j)
	// 		fmt.Printf("%p, %p\n", &i, &j)
	// 	}()
	// }

	// f := fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }

	fmt.Printf("---- %s\n", " -----")

	// s := "I am learning Go!"
	// ss := strings.Fields(s)
	// fmt.Printf("%q\n", ss)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v address=%p\n", len(x), cap(x), x, &x)
}
