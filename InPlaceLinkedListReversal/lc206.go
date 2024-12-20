package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var curr, prev, next *ListNode
	curr = head
	prev = nil

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev

}
func main() {
	//construct a list and pass to reverseList
	fmt.Println(reverseList(nil))

}
