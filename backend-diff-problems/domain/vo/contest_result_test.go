package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Last_正常系(t *testing.T) {
	list := ContestResultList{
		{true, Rating{
			rating: 100,
			color:  Gray,
		}, 1000},
		{true, Rating{
			rating: 200,
			color:  Gray,
		}, 2000},
		{true, Rating{
			rating: 300,
			color:  Gray,
		}, 3000},
	}

	actual, err := list.Last()
	assert.Nil(t, err)

	expected := ContestResult{true, Rating{rating: 300, color: Gray}, 3000}
	assert.Equal(t, expected, actual)
}

func Test_ExactByRated_正常系(t *testing.T) {
	list := ContestResultList{
		{true, Rating{rating: 100, color: Gray}, 1000},
		{true, Rating{rating: 200, color: Gray}, 2000},
		{false, Rating{rating: 300, color: Gray}, 3000},
	}

	actual := list.ExactByRated()

	expected := ContestResultList{
		{true, Rating{rating: 100, color: Gray}, 1000},
		{true, Rating{rating: 200, color: Gray}, 2000},
	}
	assert.Equal(t, expected, actual)
}
