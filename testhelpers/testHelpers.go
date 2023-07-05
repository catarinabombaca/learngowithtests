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
