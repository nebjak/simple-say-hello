package greet

import "testing"

func TestSayHello(t *testing.T) {
	name := "John"
	expect := "Hello, John!"
	got := SayHello(name)

	if got != expect {
		t.Errorf("got %q expect %q", got, expect)
	}
}
