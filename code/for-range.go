package main

import (
	"fmt"
)

func main() {
	 str := "Go is a beautiful language!"
    fmt.Printf("The length of str is: %d\n", len(str))
    for pos, char := range str {
        fmt.Printf("Character on position %d is: %c \n", pos, char)
    }
}
