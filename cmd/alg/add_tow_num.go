//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
package alg

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	temp := l3
	var sum, mode int
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			sum = l2.Val + mode
		} else if l2 == nil {
			sum = l1.Val + mode
		} else {
			sum = l1.Val + l2.Val + mode
		}
		temp.Next = &ListNode{
			Val: sum % 10,
		}
		temp = temp.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		mode = sum / 10
		if mode != 0 {
			temp.Next = &ListNode{
				Val: mode,
			}
		}
	}
	return l3.Next
}
