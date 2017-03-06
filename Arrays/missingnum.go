package main
import "fmt"

func main() {
    // Setup by creating an array from 1 -> 100, and removing an element.
    nums := SliceToN(100)
    k := 10
    nums = append(nums[:k-1], nums[k:]...)

    missing := MissingNum(nums)
    fmt.Println("Missing num:", missing);
    fmt.Println("..should be:", k)
}

// Creates a slice of size n, containing the integers
// from one to n.
func SliceToN(n int) []int {
    s := make([]int, n)
    for i := 1; i <= n; i++ {
        s[i-1] = i;
    }
    return s
}

// Finds the missing number assuming the slice is of size n-1, containing
// numbers from 1->n, missing only one number.
func MissingNum(arr []int) int {
    sum := 0
    for _,num := range arr {
        sum += num;
    }
    n := len(arr) + 1
    return ((n * (n+1)) / 2) - sum;
}