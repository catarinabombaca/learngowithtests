package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, Chris."
	got := Hello("Chris")

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}

}
