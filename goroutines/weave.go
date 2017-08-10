package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	say("wait for me!")
	go say("no need to wait!")
	say("hello!")
}
