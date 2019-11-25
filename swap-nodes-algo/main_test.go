package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//    1
//   / \
//  2   3
func TestWalkInBinaryTree_0(t *testing.T) {
	root := int32(1)
	parentRoot := int32(-1)

	indexes := [][]int32{
		{2, 3},
		{-1, -1},
		{-1, -1},
	}

	expected := []int32{
		2, 1, 3,
	}

	actual := walkInBinaryTree(root, parentRoot, indexes)
	assert.Equal(t, expected, actual)
}

//       1
//      / \
//     2   3
//    / \
//   4   5
func TestWalkInBinaryTree_1(t *testing.T) {
	root := int32(1)
	parentRoot := int32(-1)

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

	actual := walkInBinaryTree(root, parentRoot, indexes)
	assert.Equal(t, expected, actual)
}

func TestGetBinaryTreeState(t *testing.T) {
	root := int32(1)
	parentRoot := int32(-1)

	indexes := [][]int32{
		{2, 3},
		{4, 5},
		{6, -1},
		{-1, 7},
		{8, 9},
		{10, 11},
		{12, 13},
		{-1, 14},
		{-1, -1},
		{15, -1},
		{16, 17},
		{-1, -1},
		{-1, -1},
		{-1, -1},
		{-1, -1},
		{-1, -1},
		{-1, -1},
	}

	expected := []int32{
		12, 7, 13, 4, 2, 14, 8, 5, 9, 1, 15, 10, 6, 16, 11, 17, 3,
	}

	actual := walkInBinaryTree(root, parentRoot, indexes)
	assert.Equal(t, expected, actual)
}

func TestSwapNodesWithNoQueries(t *testing.T) {

	indexes := [][]int32{
		{2, 3},
		{-1, -1},
		{-1, -1},
	}

	queries := []int32{}

	expected := [][]int32{
		{2, 1, 3},
	}

	actual := swapNodes(indexes, queries)
	assert.Equal(t, expected, actual)
}

func TestSwapIndexesDepth1(t *testing.T) {

	indexes := [][]int32{
		{2, 3},
		{-1, -1},
		{-1, -1},
	}

	depth := int32(1)

	expected := [][]int32{
		{3, 2},
		{-1, -1},
		{-1, -1},
	}

	actual := swapIndexes(indexes, depth)
	assert.Equal(t, expected, actual)
}

func TestSwapIndexesDepth2(t *testing.T) {

	indexes := [][]int32{
		{2, 3},
		{4, 5},
		{-1, -1},
		{-1, -1},
		{-1, -1},
	}

	depth := int32(2)

	expected := [][]int32{
		{2, 3},
		{5, 4},
		{-1, -1},
		{-1, -1},
		{-1, -1},
	}

	actual := swapIndexes(indexes, depth)
	assert.Equal(t, expected, actual)
}

//    1            1          depth [0]
//   / \   ==>    / \
//  2   3        3   2        depth [1]
func TestSwapNodesOnceDepth1(t *testing.T) {

	indexes := [][]int32{
		{2, 3},
		{-1, -1},
		{-1, -1},
	}

	queries := []int32{
		1,
	}

	expected := [][]int32{
		{3, 1, 2},
	}

	actual := swapNodes(indexes, queries)
	assert.Equal(t, expected, actual)
}

//    1            1            1          depth [0]
//   / \   ==>    / \   ==>    / \
//  2   3        3   2        2   3        depth [1]
func TestSwapNodesTwiceDepth1(t *testing.T) {

	indexes := [][]int32{
		{2, 3},
		{-1, -1},
		{-1, -1},
	}

	queries := []int32{
		1,
		1,
	}

	expected := [][]int32{
		{3, 1, 2},
		{2, 1, 3},
	}

	actual := swapNodes(indexes, queries)
	assert.Equal(t, expected, actual)
}

//       1                  1           depth [0]
//      / \                / \
//     2   3     ==>      2   3         depth [1]
//    / \                / \
//   4   5              5   4           depth [2]
func TestSwapNodesOnceDepth2(t *testing.T) {

	indexes := [][]int32{
		{2, 3},
		{4, 5},
		{-1, -1},
		{-1, -1},
		{-1, -1},
	}

	queries := []int32{
		2,
	}

	expected := [][]int32{
		{5, 2, 4, 1, 3},
	}

	actual := swapNodes(indexes, queries)
	assert.Equal(t, expected, actual)
}
