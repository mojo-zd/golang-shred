package ord

type Node struct {
	Value int
	Next  *Node
}

func AddTwoLink(first, second *Node) *Node {
	newNode := new(Node)
	head := newNode
	var temp, pre int
	for {
		if first != nil || second != nil {
			if first == nil {
				temp = second.Value
				second = second.Next
			} else if second == nil {
				temp = first.Value
				first = first.Next
			} else {
				temp = first.Value + second.Value
				first = first.Next
				second = second.Next
			}

			sum := temp + pre
			newNode.Value = sum % 10
			pre = sum / 10

			if first == nil && second == nil {
				break
			}
			newNode.Next = new(Node)
			newNode = newNode.Next
		}
	}
	return head
}
