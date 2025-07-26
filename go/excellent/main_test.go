package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOdd(10)
	if result != "Even" {
		t.Errorf("ecpected: even, actula: %s", result)
	}
}
