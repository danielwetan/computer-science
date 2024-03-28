package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Daniel")

	got := buffer.String()
	want := "Hello, Daniel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
