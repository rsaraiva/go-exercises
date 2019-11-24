package repeated_string

import (
	"strings"
)

func CountALetterCoccurrences(s string, n int64) int64 {

	var lenString int64 = int64(len(s))

	var letterAaCount int64 = int64(strings.Count(s, "a"))

	var letterAaIndexes []bool
	for _, v := range s {
		if string(v) == "a" {
			letterAaIndexes = append(letterAaIndexes, true)
			continue
		}
		letterAaIndexes = append(letterAaIndexes, false)
	}

	var necessaryRepeatsInt int64 = n / lenString
	var necessaryRepeatsMod int64 = n % lenString

	var repeatedLetterAaCount int64 = necessaryRepeatsInt * letterAaCount

	if necessaryRepeatsMod == 0 {
		return repeatedLetterAaCount
	}

	var i int64
	for i = 0; i < necessaryRepeatsMod; i++ {
		if letterAaIndexes[i] {
			repeatedLetterAaCount++
		}
	}

	return repeatedLetterAaCount
}
