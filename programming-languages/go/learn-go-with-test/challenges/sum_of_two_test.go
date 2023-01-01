package challenges

import "testing"

func TestSumOfTwo(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{10, 20, 30, 40}
		v := 42 // 40 + 2

		got := sumOfTwo(a, b, v)
		want := true

		if got != want {
			t.Errorf("got %v, but want %v", got, want)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		a := []int{0, 0, -5, 2}
		b := []int{-10, 40, -3, 9}
		v := -8 // -5 + -3

		got := sumOfTwo(a, b, v)
		want := true

		if got != want {
			t.Errorf("got %v, but want %v", got, want)
		}
	})
}
