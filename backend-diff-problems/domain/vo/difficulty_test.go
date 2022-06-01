package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MakeClipDifficulty(t *testing.T) {
	rawDifficulties := []RawDifficulty{
		{0, false},
		{100, false},
		{-100, true},
		{399, true},
		{400, true},
		{1000, true},
	}
	expected := []ClipDifficulty{
		{0, false},
		{0, false},
		{114.60191874407604, true},
		{399.0012489589841, true},
		{400, true},
		{1000, true},
	}

	for ind, diff := range rawDifficulties {
		clip := diff.MakeClipDifficulty()
		assert.Equal(t, expected[ind], clip)
	}
}
