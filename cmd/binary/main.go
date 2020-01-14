package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func main() {
	nodeG := Node{data: "g", left: nil, right: nil}
	nodeF := Node{data: "f", left: &nodeG, right: nil}
	nodeE := Node{data: "e", left: nil, right: nil}
	nodeD := Node{data: "d", left: &nodeE, right: nil}
	nodeC := Node{data: "c", left: nil, right: nil}
	nodeB := Node{data: "b", left: &nodeD, right: &nodeF}
	nodeA := Node{data: "a", left: &nodeB, right: &nodeC}
	fmt.Println("广度优先", breadthFirstSearch(&nodeA))
	fmt.Println("先序遍历", preOrderSearch(&nodeA))
	fmt.Println("中序遍历", midOrderSearch(&nodeA))
}

// 广度优先
func breadthFirstSearch(node *Node) []string {
	var result []string
	if node == nil {
		log.Warn().Msg("node can't be nil")
		return result
	}
	nodes := []Node{*node}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		result = append(result, node.data)
		if node.left != nil {
			nodes = append(nodes, *node.left)
		}
		if node.right != nil {
			nodes = append(nodes, *node.right)
		}
	}
	return result
}

type SeqStack struct {
	data [100]*Node
	top  int // 数组下标
}

// 先序遍历
func preOrderSearch(node *Node) []string {
	var result []string
	var s SeqStack
	s.top = -1
	if node == nil {
		log.Warn().Msg("node can't be nil")
		return result
	}

	for node != nil || s.top != -1 {
		for node != nil {
			result = append(result, node.data)
			s.top++
			s.data[s.top] = node
			node = node.left
		}
		s.top--
		node = s.data[s.top+1]
		node = node.right
	}
	return result
}

// 中序遍历
func midOrderSearch(node *Node) []string {
	var result []string
	var s SeqStack
	s.top = -1

	for node != nil || s.top != -1 {
		for node != nil {
			s.top++
			s.data[s.top] = node
			node = node.left
		}
		s.top--
		node = s.data[s.top + 1]
		result = append(result, node.data)
		node = node.right
	}
	return result
}
