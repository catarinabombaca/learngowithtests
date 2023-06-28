package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		want := "Hello, Chris"
		got := Hello("Chris")

		assertCorrectMessage(t, got, want)
	})

	t.Run("return world when the name is empty", func(t *testing.T) {
		want := "Hello, world"
		got := Hello("")

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
