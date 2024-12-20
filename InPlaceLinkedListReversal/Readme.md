In Place Linked List Reversal

Use two pointers which point to the previous and current nodes.
To reverse a linked list:  ptr.next = prev
Then move prev to ptr and move ptr to the next node.  At the end of the algo
prev will point to the head of the reversed list.

def reverse_linked_list(head):
prev = None
ptr = head

    while ptr:
        # Save the next node
        next_node = ptr.next

        # Reverse the current node's pointer
        ptr.next = prev

        # Move the pointers one step forward
        prev = ptr
        ptr = next_node

    # prev is the new head after the loop ends
    return prev

When to use it?
-Reverse a linked list in 1 pass and O(1) space
-Reverse a specific portion of a linked list
-Reverse nodes in groups of k

LeetCode Questions
206. [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/description/)

143. [Reorder List](https://leetcode.com/problems/reorder-list/description/)

25.[ Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/description/)

