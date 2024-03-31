package dota2_senate

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcases = []domains.Testcase{
	{In: "RD", Out: "Radiant"},
	{In: "RDD", Out: "Dire"},
}

func TestDota2Senate(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(string)
		output := tt.Out.(string)
		assert.Equal(t, output, predictPartyVictory(input))
	}
}
