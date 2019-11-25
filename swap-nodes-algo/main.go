package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the swapNodes function below.
 */
func swapNodes(indexes [][]int32, queries []int32) [][]int32 {

	var ret [][]int32

	if len(queries) == 0 {
		return append(ret, getBinaryTreeState(indexes))
	}

	indexes_ := indexes
	for _, v := range queries {
		indexes_ = swapIndexes(indexes_, v)
		ret = append(ret, getBinaryTreeState(indexes_))
	}

	return ret
}

func swapIndexes(indexes [][]int32, depth int32) [][]int32 {
	var ret [][]int32
	for i := int32(1); i <= int32(len(indexes)); i++ {
		if i == depth {
			ret = append(ret, []int32{indexes[i-1][1], indexes[i-1][0]})
		} else {
			ret = append(ret, indexes[i-1])
		}
	}
	return ret
}

func getBinaryTreeState(indexes [][]int32) []int32 {
	root := int32(1)
	parentRoot := int32(-1)

	return walkInBinaryTree(root, parentRoot, indexes)
}

func walkInBinaryTree(root int32, parentRoot int32, indexes [][]int32) []int32 {
	var ret []int32

	node := indexes[root-1]

	left := node[0]
	right := node[1]

	hasLeft := left != int32(-1)
	hasRight := right != int32(-1)

	if hasLeft {
		r := walkInBinaryTree(left, root, indexes)
		ret = append(ret, r...)
		ret = append(ret, root)
	}

	if hasRight {
		r := walkInBinaryTree(right, root, indexes)
		ret = append(ret, r[0])
	}

	if !hasLeft && !hasRight {
		ret = append(ret, root)
	}

	return ret
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var indexes [][]int32
	for indexesRowItr := 0; indexesRowItr < int(n); indexesRowItr++ {
		indexesRowTemp := strings.Split(readLine(reader), " ")

		var indexesRow []int32
		for _, indexesRowItem := range indexesRowTemp {
			indexesItemTemp, err := strconv.ParseInt(indexesRowItem, 10, 64)
			checkError(err)
			indexesItem := int32(indexesItemTemp)
			indexesRow = append(indexesRow, indexesItem)
		}

		if len(indexesRow) != int(2) {
			panic("Bad input")
		}

		indexes = append(indexes, indexesRow)
	}

	queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries []int32

	for queriesItr := 0; queriesItr < int(queriesCount); queriesItr++ {
		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := swapNodes(indexes, queries)

	for resultRowItr, rowItem := range result {
		for resultColumnItr, colItem := range rowItem {
			fmt.Fprintf(writer, "%d", colItem)

			if resultColumnItr != len(rowItem)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		if resultRowItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
