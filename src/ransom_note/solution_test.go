package ransom_note

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type RansomNoteIn struct {
	RansomNote string
	Magazine   string
}

var testcase = []domains.Testcase{
	{
		In: RansomNoteIn{
			RansomNote: "a",
			Magazine:   "b",
		},
		Out: false,
	},
	{
		In: RansomNoteIn{
			RansomNote: "aa",
			Magazine:   "ab",
		},
		Out: false,
	},
	{
		In: RansomNoteIn{
			RansomNote: "aa",
			Magazine:   "aab",
		},
		Out: true,
	},
	{
		In: RansomNoteIn{
			RansomNote: "bg",
			Magazine:   "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj",
		},
		Out: true,
	},
}

func TestRansomNote(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(RansomNoteIn)
		output := tt.Out.(bool)
		assert.Equal(t, output, canConstruct(input.RansomNote, input.Magazine))
	}
}
