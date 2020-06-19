package main

import (
	"bytes"
	"testing"
)

/*
3
2
1
Go!
*/

type spySleeper struct {
	calls int
}

func (ss *spySleeper) Sleep() {
	ss.calls++
}

func TestCountDown(t *testing.T) {
	buff := &bytes.Buffer{}
	fakeSleeper := &spySleeper{}
	countDown(buff, fakeSleeper)

	got := buff.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if fakeSleeper.calls != 4 {
		t.Fatalf("not enough calls to sleeper, want %d got %d", fakeSleeper.calls, 4)
	}
}
