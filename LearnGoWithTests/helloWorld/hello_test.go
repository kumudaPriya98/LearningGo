package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Kumuda")
	want := "Hello Kumuda"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}