package test

import "testing"

func TestHello(t *testing.T) {
	got := HelloMustReturnString()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
