package iterator

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("a")
	expected := "aaaa"

	if repeat != expected {
		t.Errorf("expected %q but recieved %q", expected, repeat)
	}
}
