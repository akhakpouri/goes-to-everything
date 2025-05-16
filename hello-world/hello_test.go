package learning

import "testing"

func TestHello(t *testing.T) {
	got := SayHello("world")
	want := "hello, world"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
