package challenges

import "testing"

func TestReverseNum(t *testing.T) {
	t.Run("test case 1", func (t *testing.T) {
		got := reverseNum(1234)
		want := 4321
	
		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 2", func (t *testing.T) {
		got := reverseNum(123456789)
		want := 987654321
	
		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("test case 3", func (t *testing.T) {
		got := reverseNum(110099882214)
		want := 412288990011
	
		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

}