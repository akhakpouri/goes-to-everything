package learning

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := SayHello("Ali")
		want := "hello, Ali"
		assertMessage(t, got, want)
	})
	t.Run("saying 'hello, world' when an empty string is provided", func(t *testing.T) {
		got := SayHello("")
		want := "hello, world"
		assertMessage(t, got, want)
	})
}

func assertMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
