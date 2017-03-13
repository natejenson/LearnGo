package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	inputs := []string{"aaa", "abbccc", "aaabbccccd"}

	for _, s := range inputs {
		fmt.Println(s, ":", compress(s))
	}
}

// Compresses the string using Run-length encoding.
func compress(s string) string {
	if len(s) == 0 {
		return ""
	}

	var buffer bytes.Buffer
	cur := rune(s[0])
	count := 1
	for _, char := range s[1:] {
		if char != cur {
			buffer.WriteString(strconv.Itoa(count) + string(cur))
			count = 1
			cur = char
		} else {
			count++
		}
	}

	buffer.WriteString(strconv.Itoa(count) + string(cur))
	return buffer.String()
}
