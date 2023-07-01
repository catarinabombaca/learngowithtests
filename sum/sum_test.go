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

func TestSumAll(t *testing.T) {
	t.Run("mske the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		testhelpers.AssertCorrectSlice(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2})
		want := []int{2}

		testhelpers.AssertCorrectSlice(t, got, want)
	})
}
