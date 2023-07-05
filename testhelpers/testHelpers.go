package testhelpers

import (
	"reflect"
	"testing"
)

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

func AssertCorrectSlice(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertCorrectFloat(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g, want: %g", got, want)
	}
}

func AssertError(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didnt get one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("wanted no error, but got one: %q", got)
	}
}
