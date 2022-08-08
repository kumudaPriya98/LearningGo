package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)		
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("A")
	}
}