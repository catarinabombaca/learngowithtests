package iteration

import (
	"fmt"
	"github.com/catarinabombaca/learngowithtests/testhelpers"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeats a character 6 times", func(t *testing.T) {
		want := "aaaaaa"
		got := Repeat("a", 6)

		testhelpers.AssertCorrectString(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

func ExampleRepeat() {
	chars := Repeat("a", 5)
	fmt.Println(chars)
	// Output: aaaaa
}
