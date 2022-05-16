package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Last_正常系(t *testing.T) {
	list := ContestResultList{
		{true, 100, 1000},
		{true, 100, 2000},
		{true, 100, 3000},
	}

	actual, err := list.Last()
	assert.Nil(t, err)

	expected := ContestResult{true, 100, 3000}
	assert.Equal(t, expected, actual)
}

func Test_ExactByRated_正常系(t *testing.T) {
	list := ContestResultList{
		{true, 100, 1000},
		{true, 100, 2000},
		{false, 100, 3000},
	}

	actual := list.ExactByRated()

	expected := ContestResultList{
		{true, 100, 1000},
		{true, 100, 2000},
	}
	assert.Equal(t, expected, actual)
}
