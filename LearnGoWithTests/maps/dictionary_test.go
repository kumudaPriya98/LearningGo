package main

import "testing"

func TestMaps(t *testing.T) {
	dictionary := map[string]string{"test": "Sample Value"}

	got := Search(dictionary, "test")
	want := "Sample Value"

	if got != want {
		t.Errorf("got %q want %q\n", got, want)
	}
}
