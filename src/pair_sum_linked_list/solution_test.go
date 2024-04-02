package pair_sum_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPairSum(t *testing.T) {
	input := &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}
	output := 6
	assert.Equal(t, output, pairSum(input))
}
