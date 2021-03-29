package alg

import (
	"strings"
)

// 滑动窗口 求最大子串
// 1. 定义窗口大小为[0:0], end向右移动 end++
// 2. 当右边遇到有重复字符时, start向右移动 start+=index+1
// 3. 计算最大子串的长度, end+1 - start
// 4. 如果要返回子串提前记录并返回
func MaxSubStr(arr string) (int, string) {
	var start, end, max int
	var substr string
	for ; end < len(arr); end++ {
		idx := strings.IndexByte(arr[start:end], arr[end])
		if idx != -1 {
			start += idx + 1
		}
		sub := end + 1 - start
		if sub > max {
			substr = arr[start : end+1]
			max = sub
		}
	}
	return max, substr
}
