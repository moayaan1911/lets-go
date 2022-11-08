package main

import "fmt"

func main() {
	// printf prints a variable in the same way c++ does with %d meaning integer, %s meaning string and %f meaning float and all prints happening on same line unless used \n
	name, age := "ayaan", 20
	fmt.Printf("My name is %s",name)
	fmt.Printf("My age is %d",age)
}