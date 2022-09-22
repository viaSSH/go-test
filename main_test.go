package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	s := Calc(2, 3)

	if s != 5 {
		t.Error("Wrong result !!!!")
	}
}

func TestMyTest(t *testing.T) {
	var a int = 1
	var b int = 1
	if a != b {
		t.Error("miss match")
	}
}

func TestMulti(t *testing.T) {
	s := Multi(4, 5)

	if s != 20 {
		t.Error("Wrong Multiple")
	}
}
