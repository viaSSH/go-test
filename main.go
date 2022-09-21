package main

import (
	"fmt"
)

func calc(a int, b int) int {

	return a + b
}

func main() {

	fmt.Println("hello")

	var c = calc(1, 2)

	fmt.Println(c)

}
