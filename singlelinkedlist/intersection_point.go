package singlelinkedlist

/*
	给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
*/

func GetIntersectionNode(list1, list2 *List[int]) *Node[int] {
	if list1 == nil || list2 == nil {
		return nil
	}

	if list1.Head == nil || list2.Head == nil {
		return nil
	}

	f := list1.Head
	s := list2.Head

	// 当两个指针都经过两次切换后还没有找到交点，则说明没有交点
	switchCountF := 0
	switchCountS := 0

	for {
		if f == s {
			return f
		}

		f = f.Next
		s = s.Next

		if f == nil {
			switchCountF++
			f = list2.Head
		}

		if s == nil {
			switchCountS++
			s = list1.Head
		}

		// 如果两个指针都已经切换了两次链表，说明没有交点
		if switchCountF >= 2 && switchCountS >= 2 {
			return nil
		}
	}
}

/*
判断单链表是否有环
*/
func HasCircle(list *List[int]) bool {
	if list == nil {
		return false
	}

	if list.Head == nil {
		return false
	}

	fast := list.Head
	slow := list.Head
	var begin bool

	for {
		if fast == slow && begin {
			return true
		}
		begin = true

		if fast == nil {
			return false
		}
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next

		if slow == nil {
			return false
		}

		slow = slow.Next
	}
}
