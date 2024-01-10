// main.go
package main

import (
	"errors"
	"fmt"
)

func IntToString(num int) (string, error) {
	if num < 0 {
		return "", errors.New("negative numbers are not allowed")
	}
	return fmt.Sprintf("%d", num), nil
}

// func main() {
// 	s1, err := IntToString(15)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(s1)
// }
