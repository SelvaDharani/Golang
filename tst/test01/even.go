package main

import "fmt"

func IsEven(number int) bool {
	return number%2 == 0
}

func main() {
	value := IsEven(5)
	fmt.Println(value)
}
