package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	minBribes := calculateMinimumBribes(q)
	if minBribes == -1 {
		fmt.Println("Too chaotic")
		return
	}
	fmt.Println(minBribes)
}

// buble sort alg
func calculateMinimumBribes(q []int32) int32 {

	var bribesCount int32

	for i := len(q) - 1; i >= 0; i-- {
		currPosition := i + 1
		currValue := int(q[i])

		if currValue != currPosition {

			if int(q[i-1]) == currPosition {
				bribesCount++
				q = swap(q, i, i-1)

			} else if int(q[i-2]) == currPosition {
				bribesCount += 2
				q = swap(q, i-2, i-1)
				q = swap(q, i-1, i)

			} else {
				return -1
			}
		}
	}

	return bribesCount
}

func swap(q []int32, a, b int) []int32 {
	q[a], q[b] = q[b], q[a]
	return q
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
