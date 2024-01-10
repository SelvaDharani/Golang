package main

import "fmt"

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Calculate")
	res := Calculate(4)
	fmt.Println(res)
}
