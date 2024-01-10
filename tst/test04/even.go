// main.go
package main

import "fmt"

func IsEven(number int) bool {
	//make the logic error and run tests and observe it
	//return number%5 == 0
	return number%2 == 0
}

func main() {
	fmt.Println("IsEven function")
	value := IsEven(5)
	fmt.Println(value)
}
