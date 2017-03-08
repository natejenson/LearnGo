package main
import (
    "fmt"
    "strings"
    "bytes"
)

func main() {
    words := []string{"learning", "go", "is", "a", "blast! ❤"}
    fmt.Printf("%-8s %s\n", "Dumb:", naiveConcat(words))
    fmt.Printf("%-8s %s\n", "Faster:", concat(words))
    fmt.Printf("%-8s %s\n", "Join:", strings.Join(words, " "))
}

func concat(tokens []string) string {
    var buffer bytes.Buffer
    for _, token := range tokens {
        buffer.WriteString(token + " ")
    }
    result := buffer.String()
    return result[:len(result)-1]
}

func naiveConcat(tokens []string) string {
    var result  = ""
    for _, token := range tokens {
        result += token + " "
    }
    return result[:len(result)-1]
}