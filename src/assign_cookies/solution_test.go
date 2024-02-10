package assign_cookies

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"sort"
	"testing"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j, count := 0, 0, 0

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
			j++
			count++
		} else {
			j++
		}
	}
	return count
}

type Input struct {
	g []int
	s []int
}

var testcase = []domains.Testcase{
	{
		In: Input{
			g: []int{1, 2, 3},
			s: []int{1, 1},
		},
		Out: 1,
	},
	{
		In: Input{
			g: []int{1, 2},
			s: []int{1, 2, 3},
		},
		Out: 2,
	},
	{
		In: Input{
			g: []int{10, 9, 8, 7},
			s: []int{5, 6, 7, 8},
		},
		Out: 2,
	},
}

func TestAssignCookies(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(Input)
		output := tt.Out.(int)
		assert.Equal(t, output, findContentChildren(input.g, input.s))
	}
}
