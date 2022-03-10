// 17ms (38.87%), 4.6MB (58.63%), 2022/3/10 13:05
func addTwoNumbers(augend *ListNode, addend *ListNode) *ListNode {
    var(
        headPtr *ListNode = nil
        tailPtr *ListNode = nil
        newNode *ListNode = nil
        countUp int       = 0
    )
    
    for augend != nil && addend != nil {
        newNode = &ListNode{
            Val: (augend.Val + addend.Val + countUp) % 10,
            Next: nil,
        }
        countUp = (augend.Val + addend.Val + countUp) / 10

        
        if headPtr == nil {
            headPtr = newNode
        } else {
            tailPtr.Next = newNode
        }
        tailPtr = newNode
        
        augend = augend.Next
        addend = addend.Next
    }
    
    for augend != nil {
        newNode = &ListNode{
            Val: (augend.Val + countUp) % 10,
            Next: nil,
        }
        countUp = (augend.Val + countUp) / 10
        
        tailPtr.Next = newNode
        tailPtr = newNode
        augend = augend.Next
    }
    
    for addend != nil {
        newNode = &ListNode{
            Val: (addend.Val + countUp) % 10,
            Next: nil,
        }
        countUp = (addend.Val + countUp) / 10
        
        tailPtr.Next = newNode
        tailPtr = newNode
        addend = addend.Next
    }
    
    if countUp != 0 {
        newNode = &ListNode{
            Val: 1,
            Next: nil,
        }
        
        countUp = 0
        tailPtr.Next = newNode
        tailPtr = newNode
    }
    
    return headPtr
}
