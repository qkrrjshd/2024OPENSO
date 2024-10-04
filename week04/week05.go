package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var f float32 = 4.30 // f := 4.30
	var i int = 55

	fmt.Println(reflect.TypeOF(f), reflect.TypeOf(i))
	fmt.Printf("%s\n", strings.Title("gunhong"))
	fmt.Println(math.Ceil(3.99))

	fmt.Printf("value i : %d\n", i)
	fmt.Print("value i : ", i, "\n")
	fmt.Println("value i : ", i)
}
