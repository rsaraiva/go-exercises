package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicArray_input0(t *testing.T) {

	n := int32(2)

	queries := [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}

	expected := []int32{
		7, 3,
	}

	actual := dynamicArray(n, queries)
	assert.Equal(t, expected, actual)
}
