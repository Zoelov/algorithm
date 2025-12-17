package singlelinkedlist

/*
翻转单链表
example: 1->2->3->4->5->NULL
reverse: 5->4->3->2->1->NULL
*/

func ReverseList(head *Node[int]) *Node[int] {
	if head == nil || head.Next == nil {
		return head
	}

	
	var pre *Node[int]
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	
	return pre
}
