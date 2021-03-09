package alg

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func InsertHead() *Node {
	head := new(Node)
	head.Data = 0
	for i := 1; i <= 10; i++ {
		node := &Node{Data: i}
		node.Next = head
		head = node
	}
	ShowLink(head)
	return head
}

func InsertTail() {
	tail := new(Node)
	head := tail
	for i := 1; i <= 10; i++ {
		node := &Node{Data: i}
		tail.Next = node
		tail = node
	}
	ShowLink(head)
}

func ShowLink(node *Node) {
	if node == nil {
		return
	}
	for {
		fmt.Println("data:", node.Data)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}

func ReverseNode(node *Node) *Node {
	pre := &Node{}
	if node == nil {
		return pre
	}

	pre.Data = node.Data
	head := pre
	for node.Next != nil {
		head = &Node{Data: node.Next.Data}
		head.Next = pre
		pre = head
		node = node.Next
	}

	return head
}
