package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//    1
//   / \
//  2   3
func TestBinaryPrintState_0(t *testing.T) {
	root := int32(1)
	parentRoot := int32(-1)
	firstLevel := int32(1)

	indexes := [][]int32{
		{2, 3},
		{-1, -1},
		{-1, -1},
	}

	expected := []int32{
		2, 1, 3,
	}

	actual := printState(root, parentRoot, firstLevel, indexes)
	assert.Equal(t, expected, actual)
}

//       1
//      / \
//     2   3
//    / \
//   4   5
func TestBinaryPrintState_1(t *testing.T) {
	root := int32(1)
	parentRoot := int32(-1)
	firstLevel := int32(1)

	indexes := [][]int32{
		{2, 3},
		{4, 5},
		{-1, -1},
		{-1, -1},
		{-1, -1},
	}

	expected := []int32{
		4, 2, 5, 1, 3,
	}

	actual := printState(root, parentRoot, firstLevel, indexes)
	assert.Equal(t, expected, actual)
}
