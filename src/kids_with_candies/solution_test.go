package kids_with_candies

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type KidsWithCandiesInput struct {
	candies      []int
	extraCandies int
}

var testcase = []domains.Testcase{
	{
		In: KidsWithCandiesInput{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
		},
		Out: []bool{true, true, true, false, true},
	},
	{
		In: KidsWithCandiesInput{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
		},
		Out: []bool{true, false, false, false, false},
	},
	{
		In: KidsWithCandiesInput{
			candies:      []int{12, 1, 12},
			extraCandies: 10,
		},
		Out: []bool{true, false, true},
	},
}

func TestKidsWithCandies(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(KidsWithCandiesInput)
		output := tt.Out.([]bool)
		assert.Equal(t, output, kidsWithCandies(input.candies, input.extraCandies))
	}
}
