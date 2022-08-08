package main

import "testing"

func TestAdder(t *testing.T) {
	got := Add(3,5)
	want := 8

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
