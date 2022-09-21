package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	s := Calc(2, 3)

	if s != 5 {
		t.Error("Wrong result !!")
	}
}

func TestMyTest(t *testing.T) {
	var a int = 1
	var b int = 2
	if a != b {
		t.Error("miss match")
	}
}
