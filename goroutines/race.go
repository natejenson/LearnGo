package main

import (
	"fmt"
	"math/rand"
	"time"
)

func repeat(n int) {
	for i := 0; i < 5; i++ {
		fmt.Println(n, " #", i)
		slp := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * slp)
	}
	fmt.Println(n, "finished!")
}

func main() {
	for i := 0; i < 5; i++ {
		go repeat(i)
	}

	var inp string
	fmt.Scanln(&inp)
}
