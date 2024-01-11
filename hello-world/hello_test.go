package helloworld

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("expected '%q' but got '%q'", expected, result)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		result := Hello("Filipe", "")
		expected := "Hello, Filipe"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("saying default hello with empty string", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, world"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		result := Hello("Filipe", "Spanish")
		expected := "Hola, Filipe"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("saying hello in French", func(t *testing.T) {
		result := Hello("Filipe", "French")
		expected := "Bonjour, Filipe"
		assertCorrectMessage(t, result, expected)
	})
}
