package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

var ErrDivideByZero = fmt.Errorf("divide by zero")

// func main() {

// 	a1 := Add(2, 3)
// 	s1 := Subtract(5, 3)
// 	d1, _ := Divide(6, 2)
// 	d2, _ := Divide(6, 0)
// 	fmt.Println(a1, s1, d1, d2)
// }
