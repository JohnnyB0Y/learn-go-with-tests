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
	got := HelloGo("Jack")
	want := chineseHelloPrefix + "Jack"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	} else {
		t.Log(want)
	}
}
