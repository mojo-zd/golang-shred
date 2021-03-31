package alg

import (
	"strings"
	"testing"

	"github.com/mojo-zd/golang-shred/cmd/alg/search"

	"github.com/mojo-zd/golang-shred/cmd/alg/ord"

	"github.com/mojo-zd/golang-shred/cmd/alg/sorts"

	"github.com/rs/zerolog/log"
)

var db struct {
	Dns string
}

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
	s := "abcdefa"
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
		} else {
			if len := i - start + 1; len > max {
				max = len
			}
		}
	}
	return max
}

func TestSub(t *testing.T) {
	str := "abcabcbbacd"
	max, sub := ord.MaxSubStr(str)
	t.Log("max sub len:", max, ",sub:", sub)
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

func TestQuick(t *testing.T) {
	arr := []int{5, 24, 17, 8, 3, 78}
	//sorts.Bubble1Sort(arr)
	arr = sorts.QuickSort(arr)
	log.Info().Interface("quick sort", arr).Send()
}

func TestBubble(t *testing.T) {
	arr := []int{5, 24, 17, 8, 3, 78}
	sorts.Bubble2Sort(arr)
	log.Info().Interface("bubble sort", arr).Send()
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 4, 7, 8, 9, 12, 34}
	t.Log("find:", search.BinarySearch(arr, 34))
}

func TestInsertSort(t *testing.T) {
	arr := []int{5, 21, 19, 18, 10, 78}
	sorts.InsertSort(arr)
	log.Info().Interface("sorted", arr).Send()
}

func TestAddTwoLink(t *testing.T) {
	first := &ord.Node{Value: 2, Next: &ord.Node{Value: 4, Next: &ord.Node{Value: 3}}}
	second := &ord.Node{Value: 5, Next: &ord.Node{Value: 6, Next: &ord.Node{Value: 4}}}
	out := ord.AddTwoLink(first, second)
	for out != nil {
		t.Log("value", out.Value)
		out = out.Next
	}
}

func TestPalindrome(t *testing.T) {
	t.Log("is it palindrome", ord.IsPalindrome("A man, a plan, a canal: Panama"))
}

func TestLinkHead(t *testing.T) {
	t.Log("insert head...")
	InsertHead()
	t.Log("insert tail...")
	InsertTail()
}

func TestLinkReverse(t *testing.T) {
	node := InsertHead()
	t.Log("========")
	ShowLink(ReverseNode(node))
}
