package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2,2)
	expected := 5

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}