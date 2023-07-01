package helloworld

import (
	"github.com/catarinabombaca/learngowithtests/testhelpers"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		want := "Hello, Chris"
		got := Hello("Chris", "")

		testhelpers.AssertCorrectString(t, got, want)
	})

	t.Run("return world when the name is empty", func(t *testing.T) {
		want := "Hello, world"
		got := Hello("", "")

		testhelpers.AssertCorrectString(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		want := "Hola, Elodie"
		got := Hello("Elodie", spanish)

		testhelpers.AssertCorrectString(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		want := "Bonjour, Thomas"
		got := Hello("Thomas", french)

		testhelpers.AssertCorrectString(t, got, want)
	})
}
