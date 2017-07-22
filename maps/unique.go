package main

import (
	"fmt"
)

func main() {
	unique := "abcdefghijkl"
	notUnique := "abcdabcd"

	fmt.Printf("Has unique characters: %s - %v\n", "(empty)", hasUnique(""))
	fmt.Printf("Has unique characters: %s - %v\n", unique, hasUnique(unique))
	fmt.Printf("Has unique characters: %s - %v\n", notUnique, hasUnique(notUnique))
}

func hasUnique(str string) bool {
	m := make(map[byte]bool)

	for i := 0; i < len(str); i++ {
		c := str[i]
		if m[c] {
			return false
		}
		m[c] = true
	}
	return true
}
