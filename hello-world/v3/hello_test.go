package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloGo(t *testing.T) {
	got := HelloGo("Chris")
	want := "Hello Go, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
