package main

import (
	"fmt"
	// "os"
)

func add(num1 int, num2 int) int {
	return num1 + num2

}

func main() {
	var answer int = add(1, 2)

	fmt.Println(answer)
}
