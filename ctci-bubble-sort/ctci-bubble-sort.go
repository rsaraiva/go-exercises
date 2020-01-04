package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countSwaps function below.
func countSwaps(a []int32) {
	_, numSwaps, first, last := sort(a)

	fmt.Printf("Array is sorted in %v swaps.\n", numSwaps)
	fmt.Printf("First Element: %v\n", first)
	fmt.Printf("Last Element: %v\n", last)
}

func sort(a []int32) ([]int32, int, int32, int32) {
	isSorted := false
	numSwaps := 0

	for !isSorted {
		isSorted = true
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a = swap(a, i, i+1)
				isSorted = false
				numSwaps++
			}
		}
	}
	return a, numSwaps, first(a), last(a)
}

func swap(a []int32, x, y int) []int32 {
	a[x], a[y] = a[y], a[x]
	return a
}

func first(a []int32) int32 {
	return a[0]
}

func last(a []int32) int32 {
	return a[len(a)-1]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	countSwaps(a)
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
