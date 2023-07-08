package greet

import (
	"bytes"
	"github.com/catarinabombaca/learngowithtests/testhelpers"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	testhelpers.AssertCorrectString(t, got, want)
}
