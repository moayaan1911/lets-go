package main

import "fmt"

func main() {
	name := "ayaan"
	fmt.Println(name)
	// name:="raj" if u write this it will throw error as it has a semicolon :
	name="raj" // instead we can write like this
	fmt.Println(name)


	// CONSTANTS
	// for declaring constant we can use const and we cannot edit it anyhow
	const age=19;
	fmt.Println(age)
	// age=20 this will throw error, even though it doesnt has a semicolon :
}