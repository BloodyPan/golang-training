package main

import (
	"fmt"
	"strings"
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

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	s := "I am learning Go!"
	ss := strings.Fields(s)
	fmt.Printf("%q\n", ss)

	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	fmt.Println("0")
	defer fmt.Println("defer")
	fmt.Println("normal")
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v address=%p\n", len(x), cap(x), x, &x)
}
