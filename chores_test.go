package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	test := Adder(1, 2)
	expected := 3

	if test != expected {
		t.Errorf("expected %d, got %d", expected, test)
	}
}
