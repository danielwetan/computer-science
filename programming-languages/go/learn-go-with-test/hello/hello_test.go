package hello

import "testing" 

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		/*
		* t.Helper() is needed to tell the test suite that this method is a helper
		* by doing this, when it fails the line number reported will be in our function
		* rather than inside our test helper
		 */
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Daniel", "")
		want := "Hello, Daniel"

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Indonesian", func(t *testing.T) {
		got := Hello("Adam", "indonesian")
		want := "Halo, Adam"
		assertCorrectMessage(t, got, want)	
	})

	t.Run("in Turkish", func(t *testing.T) {
		got := Hello("Adam", "turkish")
		want := "Merhaba, Adam"
		assertCorrectMessage(t, got, want)
	})
}
