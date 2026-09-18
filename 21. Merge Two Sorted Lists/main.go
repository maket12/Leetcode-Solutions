/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

    head := &ListNode{}
	if list1.Val < list2.Val {
		head.Val = list1.Val
		if list1.Next != nil {
			list1 = list1.Next
		} else {
			list1 = nil
		}
	} else {
		head.Val = list2.Val
		if list2.Next != nil {
			list2 = list2.Next
		} else {
			list2 = nil
		}
	}

	current := head

	for list1 != nil && list2 != nil {
		new := &ListNode{}

		if list1.Val < list2.Val {
			new.Val = list1.Val
			list1 = list1.Next
		} else {
			new.Val = list2.Val
			list2 = list2.Next
		}

		current.Next = new
		current = new
	}

	if list1 != nil  {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	return head
}