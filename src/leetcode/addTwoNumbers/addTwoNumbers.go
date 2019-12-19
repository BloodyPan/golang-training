package main

import "fmt"

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	current := node
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = &ListNode{sum % 10, nil}
		current = current.Next
	}
	if carry == 1 {
		current.Next = &ListNode{1, nil}
	}
	return node.Next
}

func getListNode(num []int) *ListNode {
	node := &ListNode{}
	current := node
	for index, n := range num {
		current.Next = &ListNode{n, nil}
		current = current.Next
		fmt.Printf("index %d: %+v\n", index, n)
	}
	return node.Next
}

func main() {
	l1 := getListNode([]int{2, 4, 3})
	fmt.Printf("%+v\n", l1)
	// for true {
	// 	node := node1
	// 	fmt.Printf("%+v\n", node.Val)
	// 	if node.Next == nil {
	// 		break
	// 	}
	// 	node1 = node1.Next
	// }

	l2 := getListNode([]int{5, 6, 7})
	fmt.Printf("%+v\n", l2)

	result := addTwoNumbers(l1, l2)
	fmt.Printf("%+v\n", result)

	for true {
		node := result
		if node == nil {
			break
		}
		fmt.Printf("%+v\n", node)
		result = result.Next
	}
}
