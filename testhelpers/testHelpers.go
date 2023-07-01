package testhelpers

import "testing"

func AssertCorrectString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func AssertCorrectInteger(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
