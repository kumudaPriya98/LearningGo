package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "Sample Value"}
	assertStrings(t, dictionary.Search("test"), "Sample Value")
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q\n", got, want)
	}
}
