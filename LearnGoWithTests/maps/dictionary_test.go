package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "Sample Value"}
	assertStrings(t, Search(dictionary, "test"), "Sample Value")
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q\n", got, want)
	}
}
