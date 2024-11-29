package main

import (
	"fmt"
	"os"
)

func text(strs ...string) {
	fmt.Println(strs)
}

func main() {
	slice := os.Args[1:]
	fmt.Println(slice[1])
	fmt.Printf("%T\n", slice[1])
	slice = append(slice, "i", "n", "h", "a")
	fmt.Println(slice, len(slice))
	text("123", "456", "789")
}
