package remove_duplicate

import "algorithm/linked_list"

/*
删除排序链表中重复出现的元素，只保留原始链表中没有重复出现的数字
”。比如输入 [1,2,3,3,4,4,5]，输出 [1,2,5]；输入 [1,1,1,2,3]，输出 [2,3]​
*/

func RemoveDuplicateElement[T comparable](head *linked_list.Node[T]) *linked_list.Node[T] {
	var zero T
	dummpy := &linked_list.Node[T]{Data: zero, Next: head}
	pre := dummpy
	current := head

	for current != nil {
		hasDuplicate := false

		for current.Next != nil && current.Data == current.Next.Data {
			current = current.Next
			hasDuplicate = true
		}

		if hasDuplicate {
			pre.Next = current.Next
		} else {
			pre = current
		}
		current = current.Next
	}

	return dummpy.Next
}
