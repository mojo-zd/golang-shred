package alg

import (
	"testing"
)

// 两数之和
func TestSum(t *testing.T) {
	var nums = []int{3, 2, 4}
	var target = 6
	t.Log(twoSum(nums, target))
}

// 两数相加
func TestAddTwoNum(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 9,
			}},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	l3 := addTwoNumber(l1, l2)
	for {
		if l3 == nil {
			return
		}
		t.Log("val:", l3.Val)
		l3 = l3.Next
	}
}

// 字符串最长子字符串
func TestSubstring(t *testing.T) {
	//s := "abcdabcebb"
	//t.Log(s[0:1])
	t.Log(lengthOfLongestSubstring("abcdabcebb"))
}

func TestMaxArea(t *testing.T) {
	var arr = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea(arr)
}
