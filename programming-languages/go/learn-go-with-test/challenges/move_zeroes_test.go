package challenges

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		input := []int{0, 1, 0, 3, 12}
		got := moveZeroes(input)
		want := []int{1, 3, 12, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		input := []int{0, 0, 0, 0, 1}
		got := moveZeroes(input)
		want := []int{1, 0, 0, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		input := []int{1, 0, 0, 2, 12}
		got := moveZeroes(input)
		want := []int{1, 2, 12, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 4", func(t *testing.T) {
		input := []int{1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0, 9}
		got := moveZeroes(input)
		want := []int{1, 9, 8, 4, 2, 7, 6, 9, 0, 0, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 5", func(t *testing.T) {
		input := []int{0}
		got := moveZeroes(input)
		want := []int{0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}
