package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	rand.Seed(time.Now().Unix())
	tree := genTree([]int{8, 3, 9, 4, 1})
	level := maxDepth(tree)
	fmt.Println("level of binary tree", level)
	slice := outOfOder(4)
	fmt.Println(slice)
}

// genTree 生成二叉树
func genTree(elements []int) *TreeNode {
	var root *TreeNode
	if len(elements) == 0 {
		return root
	}

	root = &TreeNode{Val: elements[0]}
	elements = elements[1:]
	for _, ele := range elements {
		node := root
		for node != nil {
			// 存入左子树
			if ele < node.Val {
				if node.Left == nil {
					node.Left = &TreeNode{Val: ele}
					break
				}
				node = node.Left
			} else {
				// 存入右子树
				if node.Right == nil {
					node.Right = &TreeNode{Val: ele}
					break
				}
				node = node.Right
			}
		}
	}
	return root
}

// maxDepth 获取二叉树的最大深度
func maxDepth(root *TreeNode) int {
	var result int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 { // 遍历当前层节点并把下一层节点存储queue中
			parent := queue[0]
			queue = queue[1:]
			if parent.Left != nil {
				queue = append(queue, parent.Left)
			}
			if parent.Right != nil {
				queue = append(queue, parent.Right)
			}
			size--
		}
		result++
	}
	return result
}

// outOfOder 生成指定范围内的乱序数组
func outOfOder(n int) []int {
	slice := []int{}
	for len(slice) < n {
		out := rand.Intn(n)
		if !existEle(slice, out) {
			slice = append(slice, out)
		}
	}
	return slice
}

// existEle 检查元素是否存在
func existEle(s []int, ele int) bool {
	for _, val := range s {
		if val == ele {
			return true
		}
	}
	return false
}
