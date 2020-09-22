package alg

import (
	"strings"
)

// pwwkew
func lengthOfLongestSubstring(s string) int {
	//freq := make([]int, 128)
	//var start, end int
	//for i := 0; i < len(s); i++ {
	//	val, ok := freq[s[i]]
	//	if !ok {
	//		freq[s[i]] = i // 记录每个byte的index
	//		end++
	//		continue
	//	}
	//	delete(freq, s[i])
	//	if val < i {
	//		start = val + 1
	//	}
	//	end = i + 1
	//}
	//return end - start
	return 0
}

//"aba"
//end = 1
//"a"
//end = 2
//start =1
//end = 3

func lengthOfLongestSubstrings(s string) int {
	m := map[byte]int{}
	n := len(s)
	var rk, ans = 0, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		for rk < n && m[s[rk]] == 0 {
			m[s[rk]]++
			rk++
		}
		ans = max(ans, rk-i)
	}

	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxsub(s string) int {
	if len(s) == 0 {
		return 0
	}

	var max, start int
	for i := 1; i <= len(s); i++ {
		index := strings.Index(s[start:i-1], string(s[i-1]))
		// 重复了 移动窗口
		if index >= 0 {
			start += index + 1
		}

		if len := i - start; max < len {
			max = len
		}
	}
	return max
}
