package main

import (
	"fmt"
)

func Calc(a int, b int) int {

	return a + b
}

func main() {

	fmt.Println("hello")

	var c = Calc(1, 2)

	fmt.Println(c)

}
