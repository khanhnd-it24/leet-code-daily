package min_max_difference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMaxDifference(t *testing.T) {
	type testcase struct {
		arg  int
		want int
	}

	var testcases = []testcase{
		{arg: 11891, want: 99009},
		{arg: 456, want: 900},
		{arg: 90, want: 99},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.want, minMaxDifference(tc.arg))
	}
}
