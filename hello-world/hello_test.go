package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Filipe")
	expected := "Hello, Filipe!"

	if result != expected {
		t.Errorf("Result '%s', expected '%s'", result, expected)
	}
}
