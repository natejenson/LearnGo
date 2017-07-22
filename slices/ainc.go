package main

import (
	"fmt"
)

func main() {
	tests := [][]int{}
	tests = append(tests, []int{1, 2, 3, 4, 5})
	tests = append(tests, []int{1, 2, 3, 99, 4})
	tests = append(tests, []int{1, 2, 3, 4, 0})
	tests = append(tests, []int{1, 2, 0, 99, 4})
	tests = append(tests, []int{10, 2, 3, 4, 5})
	tests = append(tests, []int{10, 2, 3, 0, 5})


	for _, test := range tests {
		fmt.Printf("%v : %v\n", test, isAlmostIncreasing(test))
	}
}

// Determines whether one or zero array elements can be removed
// to create a strictly increasing slice of numbers.
func isAlmostIncreasing(nums []int) bool {
	hasDescreased := false
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] { // decreasing. check failure conditions
			if hasDescreased {
				return false
			}
			if !canRemove(nums, i) && !canRemove(nums, i-1) {
				return false
			}
			hasDescreased = true
		}
	}
	return true
}

func canRemove(nums []int, i int) bool {
	if i == len(nums)-1 || i == 0 { // can always remove first or the last element
		return true
	}
	if nums[i-1] < nums[i+1] { // normal case. check correct ordering after removing.
		return true
	}
	return false
}
