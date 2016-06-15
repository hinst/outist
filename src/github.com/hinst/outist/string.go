package outist

import (
	"sort"
	"strconv"
	"strings"
)

const LatinString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type TStringSortingArray []string

func BoolToStr(value bool) string {
	if value {
		return "true"
	} else {
		return "false"
	}
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func UInt64ToStr(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Add0ToStr(s string, length int) string {
	//TODO: improve performance here
	for len(s) < length {
		s = "0" + s
	}
	return s
}

func CheckStringArrayContains(array []string, value string) bool {
	var result = false
	for _, item := range array {
		if item == value {
			result = true
			break
		}
	}
	return result
}

func CheckIfLatinString(s string) bool {
	var result = true
	for _, character := range s {
		var isLatinCharacter = strings.ContainsRune(LatinString, character)
		if false == isLatinCharacter {
			result = false
			break
		}
	}
	return result
}

func ClipString(s string, length int) string {
	if len(s) <= length {
		return s
	} else {
		return s[:length]
	}
}

func (array TStringSortingArray) Len() int {
	return len(array)
}

func (array TStringSortingArray) Swap(a, b int) {
	array[a], array[b] = array[b], array[a]
}

func (array TStringSortingArray) Less(a, b int) bool {
	return array[a] < array[b]
}

func SortStringArray(array []string) {
	sort.Sort(TStringSortingArray(array))
}

func FloatToStr(x float32) string {
	return strconv.FormatFloat(float64(x), 'f', 2, 32)
}
