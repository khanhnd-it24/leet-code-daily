package odd_even_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNimGame(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	output := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}}}}
	assert.Equal(t, output, oddEvenList(input))
}
