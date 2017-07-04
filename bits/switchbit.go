package main

import (
	"fmt"
)

func main() {
	tests := []int{2, 4, 6, 8, 10, 12}
	positions := []uint{0, 1, 2, 3, 2, 3}
	for pos, num := range tests {
		fmt.Printf("Switching bit %v in %b. Result: %b\n", positions[pos], num, toggleBit(num, positions[pos]))
	}
}

// Toggle the bit in the nth positon (from the right, base zero) of the given number.
func toggleBit(num int, n uint) int {
	return num ^ 1<<n
}
