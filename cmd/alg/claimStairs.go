package alg

// 爬楼梯问题 转化为递归 下面使用动态规划做解
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
func climbStairs(n int) int {
	var result, first int
	second := 1
	for i := 1; i <= n; i++ {
		result = first + second
		first = second
		second = result
	}
	return result
}
