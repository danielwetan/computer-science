package challenges

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {
		input := []int{1, 9, 8, 4, 10, 5, 2, 7, 15, 6, 16, 9}
		got := bubbleSort(input)
		want := []int{1, 2, 4, 5, 6, 7, 8, 9, 9, 10, 15, 16}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}
