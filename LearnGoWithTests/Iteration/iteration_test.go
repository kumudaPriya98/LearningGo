package iteration

import (
	"strings"
	"testing"
)

const testChar = "a"

func TestRepeat(t *testing.T) {
	got := Repeat(testChar, 8)
	want := strings.Repeat(testChar, 8)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(testChar, 5)
	}
}
