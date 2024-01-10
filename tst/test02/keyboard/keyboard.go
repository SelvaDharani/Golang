package main

import "fmt"

func keyboard() string {
	return fmt.Sprintf("Keys are working")
}

func main() {
	result := keyboard()
	fmt.Println(result)
}
