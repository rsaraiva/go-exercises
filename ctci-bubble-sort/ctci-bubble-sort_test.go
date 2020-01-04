package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort_1(t *testing.T) {
	input := []int32{3, 2, 1}
	expectedArray := []int32{1, 2, 3}
	sortedArray, numSwaps, first, last := sort(input)

	assert.Equal(t, expectedArray, sortedArray)
	assert.Equal(t, 3, numSwaps)
	assert.Equal(t, int32(1), first)
	assert.Equal(t, int32(3), last)
}

func TestSort_2(t *testing.T) {
	input := []int32{1}
	expectedArray := []int32{1}
	sortedArray, numSwaps, first, last := sort(input)

	assert.Equal(t, expectedArray, sortedArray)
	assert.Equal(t, 0, numSwaps)
	assert.Equal(t, int32(1), first)
	assert.Equal(t, int32(1), last)
}
