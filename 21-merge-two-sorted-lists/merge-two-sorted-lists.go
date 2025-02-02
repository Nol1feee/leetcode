func mergeTwoLists(first *ListNode, second *ListNode) *ListNode {
    dummy := &ListNode{}
	current := dummy

	for first != nil && second != nil {
        if first.Val > second.Val {
            current.Next = second
            second = second.Next
        } else {
            current.Next = first
            first = first.Next
        }

        current = current.Next
	}

    if first != nil {
        current.Next = first
    } else {
        current.Next = second
    }

	return dummy.Next
}
