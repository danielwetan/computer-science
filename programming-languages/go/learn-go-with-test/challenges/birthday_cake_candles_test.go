package challenges

import "testing"

func TestBirthdayCakeCandles(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		input := []int{4, 4, 1, 3}
		got := birthdayCakeCandles(input)
		want := 2

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		input := []int{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}
		got := birthdayCakeCandles(input)
		want := 5

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}


