package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Last_正常系(t *testing.T) {
	a, b, c := 100, 200, 300
	list := ContestResultList{
		{true, Rating{
			rating: &a,
			color:  Gray,
		}, 1000},
		{true, Rating{
			rating: &b,
			color:  Gray,
		}, 2000},
		{true, Rating{
			rating: &c,
			color:  Gray,
		}, 3000},
	}

	actual, err := list.Last()
	assert.Nil(t, err)

	expected := ContestResult{true, Rating{rating: &c, color: Gray}, 3000}
	assert.Equal(t, expected, actual)
}

func Test_ExactByRated_正常系(t *testing.T) {
	a, b, c := 100, 200, 300
	list := ContestResultList{
		{true, Rating{rating: &a, color: Gray}, 1000},
		{true, Rating{rating: &b, color: Gray}, 2000},
		{false, Rating{rating: &c, color: Gray}, 3000},
	}

	actual := list.ExactByRated()

	expected := ContestResultList{
		{true, Rating{rating: &a, color: Gray}, 1000},
		{true, Rating{rating: &b, color: Gray}, 2000},
	}
	assert.Equal(t, expected, actual)
}
