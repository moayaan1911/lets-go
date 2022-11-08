package main

import "fmt"

func main() {
	// we can check the data type of any variable using %T with a printf
	name, age, height := "ayaan", 20, 5.10
	fmt.Printf("%T %T %T",name,age,height)
}