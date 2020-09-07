package alg

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var max int
	i := 0
	j := len(height) - 1

	for ; i < j; {
		min, right := minAndMoveRight(i, j, height)
		area := min * (j - i)
		if max < area {
			max = area
		}

		if right {
			i++
		} else {
			j--
		}
	}

	return max
}

// 返回最小值 并判断移动方向
// 向右为true
func minAndMoveRight(i, j int, height []int) (int, bool) {
	if height[i] > height[j] {
		return height[j], false
	}

	return height[i], true
}
