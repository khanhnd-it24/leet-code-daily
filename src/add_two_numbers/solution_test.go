package add_two_numbers

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}
		if tmp.Val >= 10 {
			tmp.Val -= 10
			tmp.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next
	}
	return result
}

type TwoNumbers struct {
	L1 *ListNode
	L2 *ListNode
}

var testcase = []domains.Testcase{
	{
		In: TwoNumbers{
			L1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			L2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
		Out: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 8,
				},
			},
		},
	},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(TwoNumbers)
		output := tt.Out.(*ListNode)
		assert.Equal(t, output, addTwoNumbers(input.L1, input.L2))
	}
}
