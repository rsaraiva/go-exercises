package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {

	input := []int32{0, 1, 2, 3, 4}
	returned := swap(input, 2, 3)
	expected := []int32{0, 1, 3, 2, 4}

	assert.Equal(t, expected, returned)
}

func TestCase1(t *testing.T) {

	input := []int32{2, 1, 5, 3, 4}
	returned := calculateMinimumBribes(input)
	expected := int32(3)

	if returned != expected {
		t.Errorf("returned: %v, expected: %v", returned, expected)
	}
}

func TestCase2(t *testing.T) {

	input := []int32{1, 2, 5, 3, 7, 8, 6, 4}
	returned := calculateMinimumBribes(input)
	expected := int32(7)

	if returned != expected {
		t.Errorf("returned: %v, expected: %v", returned, expected)
	}
}
