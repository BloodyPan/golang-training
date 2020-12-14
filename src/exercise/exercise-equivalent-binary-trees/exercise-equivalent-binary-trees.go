package main

import (
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	t1Chan := make(chan int)
	t2Chan := make(chan int)

	go Walk(t1, t1Chan)
	go Walk(t2, t2Chan)

	for i := 0; i < 10; i++ {
		v1 := <-t1Chan
		v2 := <-t2Chan

		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	Same(tree.New(1), tree.New(1))
	Same(tree.New(1), tree.New(2))
}
