package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Filipe")

    got := buffer.String()
    want := "Hello, Filipe"

    if got != want {
		t.Errorf("got %q but want %q", got, want)
	}

}