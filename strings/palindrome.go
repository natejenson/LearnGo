package main

import (
	"fmt"
)

func main() {
	words := []string{"racecar", "nope", "radar", "kayak", "LeVEl"}

	for _, word := range words {
		fmt.Printf("isPalindrome(\"%s\") = %v\n", word, isPalindrome(word))
	}
}

func isPalindrome(w string) bool {
	n := len(w)
	for i := 0; i <= n/2; i++ {
		if w[i] != w[n-i-1] {
			return false
		}
	}
	return true
}
