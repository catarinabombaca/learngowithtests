package integers

import (
	"fmt"
	"github.com/catarinabombaca/learngowithtests/testhelpers"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	testhelpers.AssertCorrectInteger(t, got, want)
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	//Output: 3
}
