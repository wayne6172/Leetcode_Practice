package has_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// 9ms (68.75%) 4.4 MB (88.51%), O(n), 2022/3/8 22:19
// 此題空間複雜度可達O(n)
func hasCycle(head *ListNode) bool {
	var (
		normalPtr *ListNode = head
		fastPtr   *ListNode = head
	)

	for fastPtr != nil && fastPtr.Next != nil {
		normalPtr = normalPtr.Next
		fastPtr = fastPtr.Next.Next

		if normalPtr == fastPtr {
			break
		}
	}

	if fastPtr == nil || fastPtr.Next == nil {
		return false
	} else {
		return true
	}
}
