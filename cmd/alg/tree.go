package alg

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 获取二叉树的层级
func maxDepth(root *TreeNode)  int{
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