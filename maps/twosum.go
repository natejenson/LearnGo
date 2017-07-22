package main

import (
	"fmt"
)

func main() {
	a := []int{2, 4, 6, 3, -5, 6, -3}
	target := 0
	x, y := twoSum(a, target)
	fmt.Printf("2 Sum (target %d) at indices %d & %d", target, x, y)
}

// Finds the indices of two numbers in the given slice that
// sum to the target number. Returns a pair of -1's if no
// such sum exists.
func twoSum(nums []int, target int) (int, int) {
	m := make(map[int]int)
	for i, val := range nums {
		m[val] = i
	}
	for i, val := range nums {
		if j, ok := m[target-val]; ok && i != j {
			return i, j
		}
	}
	// No sum found
	return -1, -1
}
