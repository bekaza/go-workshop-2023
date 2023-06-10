package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world 2023"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWithMessage(t *testing.T) {
	got := HelloWithMessage("test")
	want := "test"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWithMessagePP(t *testing.T) {
	got := HelloWithMessage("ajasd")
	want := "ppp"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
