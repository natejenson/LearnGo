package main

import (
	"fmt"
	"strings"
)

func main() {
	tests := []string{
		"racecar",
		"natejenson",
		"kkaya",
		"RaCecaR"}

	for _, t := range tests {
		fmt.Printf("Can be palindrome: %s - %v\n", t, canBePalindrome(t))
	}
}

// Returns true if some permutation of the given string
// is a palindrome, ignoring case.
func canBePalindrome(str string) bool {
	str = strings.ToLower(str)
	counts := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		counts[str[i]]++
	}

	numOdd := 0
	for _, v := range counts {
		if v%2 == 1 {
			numOdd++
			if numOdd > 1 {
				return false
			}
		}
	}
	return true
}
