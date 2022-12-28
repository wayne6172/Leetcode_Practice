package deleteDuplicates

// 4 ms (72.20%), 3.1 MB (72.42%) 2022/3/10 12:42
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var (
		frontPtr *ListNode = nil
		nowPtr   *ListNode = head
		nextPtr  *ListNode = head.Next
	)

	for nextPtr != nil {
		if nowPtr.Val == nextPtr.Val {
			for nextPtr != nil && nowPtr.Val == nextPtr.Val {
				nowPtr = nextPtr
				nextPtr = nextPtr.Next
			}

			if frontPtr == nil {
				head = nextPtr
			} else {
				frontPtr.Next = nextPtr
			}
		} else {
			frontPtr = nowPtr
		}

		if nextPtr != nil {
			nowPtr = nextPtr
			nextPtr = nextPtr.Next
		}
	}

	return head
}
