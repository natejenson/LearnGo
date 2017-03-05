package main

import "fmt"

func main() {
    arr := [5]int{1,2,3,4,5}
    for i := len(arr) - 1; i >= 0; i-- {
        fmt.Println(arr[i])
    }
}