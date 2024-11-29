package main

import "fmt"

func main() {
	ages := make(map[string]int)

	var name string
	var age int

	for {
		fmt.Print("What's your name? (exit with 'q'): ")
		fmt.Scanln(&name)
		if name == "q" {
			break
		}

		fmt.Print("Your age? ")
		fmt.Scanln(&age)

		ages[name] = age
	}

	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
