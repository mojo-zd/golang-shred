package ord

import (
	"unicode"
)

// 查找回文子串
func IsPalindrome(str string) bool {
	var ss []rune
	result := true
	for _, s := range str {
		if IsNumberOrLetter(s) {
			ss = append(ss, s)
		}
	}

	length := len(ss)
	for i := 0; i < length/2; i++ {
		if unicode.ToLower(ss[i]) != unicode.ToLower(ss[length-i-1]) {
			result = false
			break
		}
	}
	return result
}

func IsNumberOrLetter(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsNumber(s)
}
