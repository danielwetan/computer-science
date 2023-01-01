package challenges

import (
	"reflect"
	"testing"
)

func TestCompareTheTriplets(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		a := []int32{1, 2, 3}
		b := []int32{3, 2, 1}
		got := compareTheTriplets(a, b)
		want := []int32{1, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		a := []int32{1, 0, 3, 4}
		b := []int32{3, 2, 1, 9}
		got := compareTheTriplets(a, b)
		want := []int32{1, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		a := []int32{1, 0}
		b := []int32{0, 1}
		got := compareTheTriplets(a, b)
		want := []int32{1, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}
