package challenges

import "testing"

func TestDiagonalDifference(t *testing.T) {
	t.Run("test case 1", func (t *testing.T) {
		input := [][]int{
			{ 1, 2, 3},
			{ 4, 5, 6},
			{ 9, 8, 9},
		}

		got := diagonalDifference(input)
		want := 2 

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 2", func (t *testing.T) {
		input := [][]int{
			{ 1, 2, 3, 5},
			{ 4, 5, 6, 1},
			{ 9, 8, 9, 2},
			{ 1, 2, 3, 5},
		}

		got := diagonalDifference(input)
		want := 0

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 3", func (t *testing.T) {
		input := [][]int{
			{ 6, 8},
			{-6, 9},
		}

		got := diagonalDifference(input)
		want := 13

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}
