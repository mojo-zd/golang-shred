package alg

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 获取二叉树的层级 深度优先
func maxDepth(root *TreeNode) int {
	var result int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue) // 获取当前层的元素总数
		// 移除当前层元素并获取当前层所有元素的子节点
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
		result++
	}
	return result
}

// 二叉树的前序遍历  每次将root压入栈顶使用以后再分别压入root的右孩子和左孩子
func preOrder(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		len := len(stack)
		last := stack[len-1]
		stack = stack[:len-1]
		result = append(result, last.Val)
		if last.Right != nil {
			stack = append(stack, last.Right)
		}
		if last.Left != nil {
			stack = append(stack, last.Left)
		}
	}
	return result
}

// 中序遍历
func midOrder(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		// 压入左孩子
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		index := len(stack) - 1
		result = append(result, stack[index].Val)
		root = stack[index].Right
		stack = stack[:index]
	}
	return result
}
