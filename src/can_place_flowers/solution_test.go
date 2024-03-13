package can_place_flowers

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type PlaceFlowersInput struct {
	flowerbed []int
	n         int
}

var testcase = []domains.Testcase{
	{
		In: PlaceFlowersInput{
			flowerbed: []int{0, 0, 1, 0, 0},
			n:         1,
		},
		Out: true,
	},
	{
		In: PlaceFlowersInput{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
		},
		Out: true,
	},
	{
		In: PlaceFlowersInput{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
		},
		Out: false,
	},
}

func TestDivisorGame(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(PlaceFlowersInput)
		output := tt.Out.(bool)
		assert.Equal(t, output, canPlaceFlowers(input.flowerbed, input.n))
	}
}
