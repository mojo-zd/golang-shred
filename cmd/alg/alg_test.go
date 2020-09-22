package alg

import (
	"github.com/rs/zerolog/log"
	"strings"
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

func TestMoveZero(t *testing.T) {
	needMove := []int{0, 1, 0, 3, 12}
	MoveZero(needMove)
	log.Info().Interface("moved", needMove).Send()
}

func TestStr(t *testing.T) {
	s := "abcdddaak"
	log.Info().Int("str", lengthOfLongestSubstring1(s)).Send()
}

func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var start, max int
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index >= 0 {
			start = i
			//log.Info().Int("start", start).Send()
		} else {
			if len := i - start + 1; len > max {
				max = len
				log.Info().Interface("start", start).Interface("i", i).Send()
			}
		}
	}
	return max
}

// 树的深度变量、前序遍历
func TestTreeOfMaxDepth(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	log.Info().Int("depth", maxDepth(root)).Send()
	log.Info().Interface("pre order", preOrder(root)).Send()
	log.Info().Interface("mid order", midOrder(root)).Send()
}

func TestClaimStairs(t *testing.T) {
	log.Info().Int("claim methods", climbStairs(4)).Send()
}

func TestRep(t *testing.T) {
	log.Info().Interface("len", maxsub("")).Send()
}

func TestQuick(t *testing.T) {
	//quick := []int{6, 3, 7, 9, 4, 2}
	//quickSort(quick, 1, len(quick))
	//log.Info().Interface("out", quick).Send()
	arr := []int{5, 24, 17, 8, 3, 78}
	bubblingSort(arr)
	log.Info().Interface("sort", arr).Send()
}
