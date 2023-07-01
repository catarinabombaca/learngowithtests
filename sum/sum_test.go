package sum

import (
	"github.com/catarinabombaca/learngowithtests/testhelpers"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("give the total sum of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		testhelpers.AssertCorrectInteger(t, got, want)
	})
}
