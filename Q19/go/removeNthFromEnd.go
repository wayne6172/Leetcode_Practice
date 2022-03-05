package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

// 0 ms (100.00%), 2.1 MB (89.29%), 2022/3/5 13:06
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		topPtr   *ListNode = nil
		frontPtr *ListNode = head
		endPtr   *ListNode = head
	)

	for i := 1; i < n; i++ {
		endPtr = endPtr.Next
	}

	for endPtr.Next != nil {
		topPtr = frontPtr
		frontPtr = frontPtr.Next
		endPtr = endPtr.Next
	}

	if topPtr == nil {
		return frontPtr.Next
	} else {
		topPtr.Next = frontPtr.Next
		return head
	}
}
