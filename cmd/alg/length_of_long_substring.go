package alg

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
