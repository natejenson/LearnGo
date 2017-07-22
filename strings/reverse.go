package main

import (
	"fmt"
)

func main() {
	word := "minneapolis"
	fmt.Println(word, ":", reverse(word))
}

func reverse(s string) string {
	r := []rune(s)
	reverseRunes(r)
	return string(r)
}

func reverseRunes(r []rune) {
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}
}
