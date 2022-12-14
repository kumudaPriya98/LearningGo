package main

import "testing"

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("Add new value", func(t *testing.T) {
		err := dictionary.Add("test", "Sample Value")

		assertNoError(t, err)
		assertValidSearch(t, dictionary, "test", "Sample Value")
	})

	t.Run("Add existing value", func(t *testing.T) {
		err := dictionary.Add("test", "Sample Value2")

		assertError(t, err, ErrDuplicateKey)
		assertValidSearch(t, dictionary, "test", "Sample Value")
	})

}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "Sample Value"}

	t.Run("Search valid String", func(t *testing.T) {
		assertValidSearch(t, dictionary, "test", "Sample Value")
	})

	t.Run("Search invalid String", func(t *testing.T) {
		value, err := dictionary.Search("unknown")

		assertError(t, err, ErrNoKey)
		assertStrings(t, value, "")
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"test": "Sample Value"}

	t.Run("Update already present Key", func(t *testing.T) {
		err := dictionary.Update("test", "Sample Value2")

		assertNoError(t, err)
		assertValidSearch(t, dictionary, "test", "Sample Value2")
	})

	t.Run("Update unkown Key", func(t *testing.T) {
		err := dictionary.Update("test2", "Sample Value")

		assertError(t, err, ErrNoKey)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "Sample Value"}

	dictionary.Delete("test")

	got, err := dictionary.Search("test")
	assertError(t, err, ErrNoKey)
	assertStrings(t, got, "")
}

func assertValidSearch(t testing.TB, dictionary Dictionary, key, want string) {
	t.Helper()

	got, err := dictionary.Search(key)
	assertNoError(t, err)
	assertStrings(t, got, want)
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
