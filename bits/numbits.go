package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 5, 8, 13, 21, 34, 255}

	for _, v := range nums {
		fmt.Printf("%v : %b = %v bits\n", v, v, numBits(v))
	}
}

// Returns the number of '1' bits in the binary representation of the given integer.
func numBits(n int) int {
	res := 0
	for ; n > 0; n >>= 1 {
		res += n & 1
	}
	return res
}
