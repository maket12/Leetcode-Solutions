/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    var lengthA, lengthB int
	nodeA, nodeB := head, head

	for nodeA != nil {
		nodeA = nodeA.Next
		lengthA++
	}
	for nodeB != nil {
		nodeB = nodeB.Next
		lengthB++
	}

	for lengthA != lengthB {
		if headA == headB {
			return headA
		}

		if lengthA > lengthB {
			headA = headA.Next
			lengthA--
		} else {
			headB = headB.Next
			lengthB--
		}
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next
	}

	return nil
}