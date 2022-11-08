package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5} // we can initialize an array like this
	fmt.Println(arr)
	arr[0]=69 // we can change the element of an array like this
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ { // we can find length of the arr by using len() just like python
		fmt.Println(arr[i])
	}

}