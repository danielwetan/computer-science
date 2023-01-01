package challenges

import "testing"

func TestDistributeCandy(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		input := []int{1, 2}
		got := distributeCandy(input)
		want := 3

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		input := []int{1, 5, 2, 1}
		got := distributeCandy(input)
		want := 7

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		input := []int{1, 1, 1, 2, 2, 3, 4, 4}
		got := distributeCandy(input)
		want := 18

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

}
