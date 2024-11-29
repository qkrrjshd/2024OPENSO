package main

import "fmt"

func main() {
	var student struct {
		id   int
		name string
		gpa  float32
	}
	student.id = 202106012
	student.name = "king park"
	student.gpa = 4.5
	fmt.Println(student.gpa)
}
