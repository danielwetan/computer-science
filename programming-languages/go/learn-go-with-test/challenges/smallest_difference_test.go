package challenges

import "testing"

func TestSmallesDifference(t *testing.T) {
	t.Run("test case 1", func (t *testing.T) {
		got := smallestDifference([]int{25, 60, 40, 75, 65})
		want := 5

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})

	t.Run("test case 2", func (t *testing.T) {
		got := smallestDifference([]int{9, 5, -7, -2, 0, 1, -1})
		want := 1

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})
}