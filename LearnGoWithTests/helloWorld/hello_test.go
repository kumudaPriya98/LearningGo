package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to myself", func(t *testing.T) {
		got := Hello("Kumuda")
		want := "Hello Kumuda"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Saying Hello by default", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
