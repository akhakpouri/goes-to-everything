package iterator

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 0)
	expected := "aaaaa"

	if repeat != expected {
		t.Errorf("expected %q but recieved %q", expected, repeat)
	}
}

func BenchmarkIterator(b *testing.B) {
	for b.Loop() {
		Repeat("a", 4)
	}
}
