package main

import "testing"

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "Sample Value")

	want := "Sample Value"
	got, err := dictionary.Search("test")

	assertNoError(t, err)
	assertStrings(t, got, want)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "Sample Value"}

	t.Run("known String", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "Sample Value"

		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("unknown String", func(t *testing.T) {
		value, err := dictionary.Search("unknown")

		assertError(t, err, ErrNoKey)
		assertStrings(t, value, "")
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("expected no error, but got %q", got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Expected error but got no error")
	}

	if got != want {
		t.Errorf("got error %q, want error %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q\n", got, want)
	}
}
