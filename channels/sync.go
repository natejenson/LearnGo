package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go fakeHTTP("really cool data", c1)
	go fakeDB("GOLANG", c2)

	for i := 0; i < 2; i++ {
		select {
		case httpMsg := <-c1:
			fmt.Println("received:", httpMsg)
		case dbMsg := <-c2:
			fmt.Println("received:", dbMsg)
		}
	}
}

func fakeHTTP(msg string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- "{ msg: '" + msg + "' }"
}

func fakeDB(msg string, c chan string) {
	time.Sleep(time.Second * 1)
	c <- "SELECT * FROM " + msg + " WHERE go = 'pher'"
}
