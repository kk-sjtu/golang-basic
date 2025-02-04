
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
	n1 := n0.Next // 因为原先n0后面是有节点的，所以需要用n1来保存一下
	P.next = n1
	n0.Next = P
}

func removeItem(n0 ListNode){
	if n0.Next == nil{
		return
	}
	P := n0.Next
	n1 := P.Next
	n0.Next = n1
}