package challenges

import (
	"reflect"
	"testing"
)

func TestSmallestIntNotPresent(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		input := []int{4, 7, 0, 2, 1, 3, 5, 9, 10, 18, -1}
		got := smallestIntNotPresent(input)
		want := 6

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		input := []int{-4, -7, 0, -2, -1, -3, -5}
		got := smallestIntNotPresent(input)
		want := -6

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		input := []int{-4, -7, -3, -5, -6, -8, -9}
		got := smallestIntNotPresent(input)
		want := -2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}
