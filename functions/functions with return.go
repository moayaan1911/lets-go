package main

// functions with arguments and with return type
import "fmt"

func main() {
	sum := f1(10, 20)
	fmt.Println(sum)
}

func f1(a int, b int) int {
	return a+b
}