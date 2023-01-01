package challenges

import "testing"

func TestTooManyTabs(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {
		input := []string{"www.google.com", "www.google.com/id", "www.google.co.id", "www.google.co.id/account", "www.ruangguru.com/ruangbelajar", "roboguru.ruangguru.com", "roboguru.ruangguru.com/faq"}
		got := tooManyTabs(input)
		want := 4

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("Test case 2", func(t *testing.T) {
		input := []string{"www.google.com", "www.google.com/id", "www.google.com/id/account", "www.google.com/id/account/reset", "www.google.com/id/account/forget_password"}
		got := tooManyTabs(input)
		want := 1

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}
