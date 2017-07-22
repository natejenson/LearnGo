package main
import (
    "fmt"
    "sort"
)

func main() {
    s := []int{3,5,7,345,6,34,5,68,0}
    sort.Ints(s)
    fmt.Println("sorted:", s)
    sort.Sort(sort.Reverse(sort.IntSlice(s)))
    fmt.Println("reverse:", s)

}