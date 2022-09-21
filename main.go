package main

import (
	"fmt"
)

func Calc(a int, b int) int {

	return a + b
}

func main() {

	fmt.Println("THIS IS TEST CODE")

	var c = Calc(1, 2)

	fmt.Println(c)

	fmt.Println("this is new part")

}
