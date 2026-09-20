func findMid(head *ListNode) *ListNode {
    var prev *ListNode
    slow, fast := head, head

    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    if prev != nil {
        prev.Next = nil
    }

    return slow
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
    }

    if list1 != nil {
        current.Next = list1
    }
    if list2 != nil {
        current.Next = list2
    }

    return dummy.Next
}

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    mid := findMid(head)

    left := sortList(head)
    right := sortList(mid)

    return merge(left, right)
}