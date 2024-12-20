package SlowAndFastPointers

import (
	"fmt"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{Val: -1, Next: head}

	prev, cur := temp, temp

	for cur.Next != nil {
		if n <= 0 {
			prev = prev.Next
		}

		cur = cur.Next
		n--
	}

	target := prev.Next
	prev.Next = target.Next

	return temp.Next
}

func main() {
	fmt.Println(removeNthFromEnd(nil, 7))

}
