package helloworld

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' but want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Filipe", "")
		want := "Hello, Filipe"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying default hello with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Filipe", "Spanish")
		want := "Hola, Filipe"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Filipe", "French")
		want := "Bonjour, Filipe"
		assertCorrectMessage(t, got, want)
	})
}
