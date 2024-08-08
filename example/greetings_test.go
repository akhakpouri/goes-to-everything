package greetings

import (
	"regexp"
	"testing"
)

// this method calls greetings.SayHello method with name, checking for valid retgurn value
func TestSayHello(t *testing.T) {
	name := "Jon"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := SayHello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`SayHello("") = %q, %v, want "", error`, msg, err)
	}
}

// this method calls greetings.SayHello method with emopty string
func TestEmptyHello(t *testing.T) {
	msg, err := SayHello("")
	if msg != "" || err == nil {
		t.Fatalf(`SayHello("") = %q, %v, want "", error`, msg, err)
	}
}
