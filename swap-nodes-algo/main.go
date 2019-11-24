package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	parent int32
	left   int32
	right  int32
}

/*
 * Complete the swapNodes function below.
 */
func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	return nil
}

func printState(root int32, parentRoot int32, level int32, indexes [][]int32) []int32 {
	var ret []int32

	node := indexes[root-1]

	level++

	left := node[0]
	right := node[1]

	hasLeft := left != int32(-1)
	hasRight := right != int32(-1)

	if hasLeft {
		r := printState(left, root, level, indexes)
		ret = append(ret, r...)
		ret = append(ret, root)
	}

	if hasRight {
		r := printState(right, root, level, indexes)
		ret = append(ret, r[0])
	}

	if !hasLeft && !hasRight {
		ret = append(ret, root)
	}

	return ret
}

func debug(s ...interface{}) {
	if false {
		fmt.Println(s...)
	}
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
