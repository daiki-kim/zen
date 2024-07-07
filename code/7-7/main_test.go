package main

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 5

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
