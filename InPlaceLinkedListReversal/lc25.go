package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}

	odd := head
	even := head.Next

	for (even != nil && even.Next != nil) || (odd != nil && odd.Next != nil) {
		even.Val, odd.Val = odd.Val, even.Val
		if even.Next != nil {
			even = even.Next.Next
		}
		if odd.Next != nil {
			odd = odd.Next.Next
		}
	}

	return dummy.Next
}

func main() {
	var head *ListNode
	fmt.Println(swapPairs(head))

}
