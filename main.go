package main

import (
	"fmt"
)

func Calc(a int, b int) int {

	return a + b
}

func Multi(a int, b int) int {
	return a * b
}

func main() {

	fmt.Println("THIS IS TEST CODE")

	var c = Calc(1, 2)
	var d = Multi(3, 4)

	fmt.Println(c, d)

}
