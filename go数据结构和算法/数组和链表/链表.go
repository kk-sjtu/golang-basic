
type ListNode struct{
	val int
	Next *ListNode
}

func NewListNode(val int) *ListNode{
	return &ListNode{
		val : val
		Next:nil
	}
}

n0 := NewListNode(0)
n1 := NewListNode(1)
n2 := NewListNode(2)
n3 := NewListNode(3)
n4	:= NewListNode(4)

n0.Next = n1
n1.Next = n2
n2.Next = n3
n3.Next = n4

func insertNode(n0 *ListNode,P *ListNode){
	n1 := n0.Next
	P.next = n1

}