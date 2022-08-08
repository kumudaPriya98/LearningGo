package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to myself", func(t *testing.T) {
		got := Hello("Kumuda")
		want := "Hello Kumuda"

		assertMessage(t, got, want)
	})

	t.Run("Saying Hello by default", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		assertMessage(t, got, want)
	})
}

// Both *testing.T and *testing.B accept testing.TB
func assertMessage(t testing.TB, got, want string) {
	// the following declaration tell the test suite that, the current method is a helper method
	// So, when the test fails, it prints the actual function line no, instead of current line no
	t.Helper()

	if got != want {
		t.Errorf("got %q wat %q", got, want)
	}

}
