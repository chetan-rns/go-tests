package iteration

import (
	"testing"
)

const repeatCount = 5

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"
	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}

func Repeat(c string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += c
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
