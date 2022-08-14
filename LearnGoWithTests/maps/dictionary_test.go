package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "Sample Value"}

	t.Run("known String", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "Sample Value"
		assertStrings(t, got, want)
	})

	t.Run("unknown String", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNoKey.Error()

		if err == nil {
			t.Fatal("Expected error but got no error")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q\n", got, want)
	}
}
