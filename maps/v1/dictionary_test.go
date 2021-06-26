package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got := dictionary.Search("test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestValueOfKey(t *testing.T) {
	dict := IntDictionary{"1": 1, "2": 2}

	got := dict.ValueForKey("1")
	want := 1

	assertInts(t, got, want)
	assertInts(t, dict.ValueForKey("2"), 2)
}

func assertInts(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
