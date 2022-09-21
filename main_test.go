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
