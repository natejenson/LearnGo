package main
import "fmt"

func main() {
    arr := []int{1,2,3,4,5}
    arrB := make([]int, len(arr))
    copy(arrB, arr)
    n := 3

    // Method 1: Rotate by one n times
    for i := 0; i < n; i++ {
        RotateLeftByOne(arr)
    }
    fmt.Println(arr)

    // Method 2: Use temp array to rotate
    RotateLeft(arrB, n)
    fmt.Println(arrB)
}

func RotateLeftByOne(arr []int) {
    temp := arr[0]
    for i := 0; i < len(arr) - 1; i++ {
        arr[i] = arr[i+1]
    }
    arr[len(arr)-1] = temp
}

func RotateLeft(arr []int, d int) {
    d = d % len(arr)

    // Store the first d items in a temp array
    temp := make([]int, d)
    copy(temp,arr[:d])

    // Shift original slice
    for i := 0; i < len(arr) - d; i++ {
        arr[i] = arr[i+d]
    }

    // Store back the temp elements
    for i := 0; i < len(temp); i++ {
        arr[len(arr) - d + i] = temp[i]
    }
}