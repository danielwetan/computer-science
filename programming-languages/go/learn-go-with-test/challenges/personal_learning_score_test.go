package challenges

import "testing"

func TestPersonalLearningScore(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {
		rule := []int{7, 10, 8}
		log := []int{6, 0, 0, 1, 5, 1, 1, 0}

		got := personalLearningScore(rule, log)
		want := 3

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("Test case 2", func(t *testing.T) {
		rule := []int{7, 10, 7}
		log := []int{1, 1, 1, 1, 1, 1, 1}

		got := personalLearningScore(rule, log)
		want := 0

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("Test case 3", func(t *testing.T) {
		rule := []int{3, 9, 7}
		log := []int{3, 3, 3, 4, 3, 3, 3}

		got := personalLearningScore(rule, log)
		want := 1

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}
