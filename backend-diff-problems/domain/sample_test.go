package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	a := 1
	b := 4
	expected := 5

	actual := add(a, b)
	assert.Equal(t, expected, actual)
}
