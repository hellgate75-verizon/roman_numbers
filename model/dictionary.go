package model

import (
	"fmt"
	"sort"
	"strings"
)

var dictionaryLength int
var keys = make([]int64, 0)
var dictionary = make(map[int64]rune)

var NumberConversionLimit = int64(3000)

// Initialize the dictionary
func Init() {
	if len(dictionary) > 0 {
		return
	}
	dictionary[1] = 'I'
	dictionary[5] = 'V'
	dictionary[10] = 'X'
	dictionary[50] = 'L'
	dictionary[100] = 'C'
	dictionary[500] = 'D'
	dictionary[1000] = 'M'
	dictionaryLength = len(dictionary)
	orderKeys()
}

// Check the input pre-conditions
func CheckNumber(i int64) bool {
	return i > 0 && i <= NumberConversionLimit
}

func orderKeys() {
	for k, _ := range dictionary {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(Int64Slice(keys)))
}

// Convert number to Roman Numbering format
// Negative, less than one or more then maxNumber
// values will returns empty output
func Convert(n int64) string {
	if ! CheckNumber(n) {
		return ""
	}
	Init()
	var output = ""
	for _, k := range keys {
		v := dictionary[k]
		if n >= k {
			// Key is matching the key
			var val int64 = n / k
			n -= val * k
			repeat := int(val)
			output = strings.Repeat(fmt.Sprintf("%c", v), repeat) + output
		}
	}
	return output
}

