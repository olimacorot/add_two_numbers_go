package add_two_numbers 

type ListNode struct {
    Val int
    Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    var result *ListNode = &ListNode{0, nil}
    current := result 
    carry   := 0

    for l1 != nil || l2 != nil || carry != 0 {
        x := 0 
        if l1 != nil {
            x = l1.Val
        }
        y := 0
        if l2 != nil {
            y = l2.Val
        }
        sum := x + y + carry
        carry = sum / 10

        current.Next = &ListNode{sum % 10, nil}
        current = current.Next
        
        if (l1 != nil) {
            l1 = l1.Next
        }
        if (l2 != nil) {
            l2 = l2.Next
        }
    }

    return result.Next
}