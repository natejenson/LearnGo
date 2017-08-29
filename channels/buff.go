package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := 3

	c := make(chan string, s)

	reader := bufio.NewReader(os.Stdin)
	msg := ""
	for ok := true; ok; ok = tryPush(msg, c) {
		fmt.Print("Enter message: ")
		msg, _ = reader.ReadString('\n')
	}
	fmt.Println("channel already full! you entered:")
	for i := 1; i <= s; i++ {
		fmt.Printf("input #%v = %v\n", i, <-c)
	}
}

func tryPush(msg string, c chan string) bool {
	select {
	case c <- msg:
		return true
	default:
		return false
	}
}
