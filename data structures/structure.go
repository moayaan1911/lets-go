package main

import "fmt"

func main() {
	// GO has struct where we can store a collection of datas of different data types together. STRUCT in Go is like struct in solidity
	type Book struct {
		name   string
		id     int
		author string
	}
	myBook := Book{"ayaan", 1, "siddiqui"} // this is how we can create an INSTANCE of a struct
	fmt.Println(myBook)
}