package main

import "fmt"

func main() {
	// hash map in go lang
	people := map[string]int{}
	people["ayaan"] = 69
	people["rohit"]=20
	fmt.Println(people)

	for key,value:=range people{ // we can also print a hash map like this
		fmt.Println(key,value)
	}

}