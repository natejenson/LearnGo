package main

import (
	"fmt"
)

func main() {
	tests := []int{1, 2, 3, 5, 8, 13, 21, 34, 255}
	for _, t := range tests {
		fmt.Printf("%v : %b = parity of %v\n", t, t, parity(t))
	}
}

// Determines the parity of the given integer, returning 1 if the number of ones in the binary
// representation of the number is odd, return 0 if even.
func parity(n int) int {
	res := 0
	for ; n > 0; n >>= 1 {
		res = res ^ n&1
	}
	return res
}
