// база, забыл о s := &ListNode{} и var s *ListNode

func reverseList(head *ListNode) *ListNode {
    var top *ListNode

    for head != nil {
        head, top, head.Next = head.Next, head, top
    }

    return top
}