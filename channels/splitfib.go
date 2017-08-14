package main

import "fmt"

func sumFib(nums []int, c chan int) {
	sum := 0
	for _, v := range nums {
		sum += fib(v)
	}
	c <- sum
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	c := make(chan int)
	go sumFib(nums[:len(nums)/2], c)
	go sumFib(nums[len(nums)/2:], c)
	x, y := <-c, <-c

	fmt.Printf("sum 1 = %v, sum 2 = %v, total: %v", x, y, x+y)
}
