package operation_x2y

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OperationX2Y(t *testing.T) {
	type testcase struct {
		name string
		x    int
		y    int
		want int
	}
	var testcases = []testcase{
		{"2", 54, 2, 4},
		{"1", 26, 1, 3},
		{"3", 25, 30, 5},
		{"4", 89, 57, 32},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.want, minimumOperationsToMakeEqual(tc.x, tc.y), tc.name)
	}
}
