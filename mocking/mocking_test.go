package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

/*
3
2
1
Go!
*/

const (
	sleep = "sleep"
	write = "write"
)

type spySleeper struct {
	calls []string
}

func (s *spySleeper) Sleep() {
	s.calls = append(s.calls, sleep)
}

func (s *spySleeper) Write(b []byte) (int, error) {
	s.calls = append(s.calls, write)
	return len(s.calls), nil
}

func TestCountDown(t *testing.T) {
	t.Run("check print", func(t *testing.T) {
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
	})
	t.Run("check sleep order", func(t *testing.T) {
		fakeSleeper := &spySleeper{}
		countDown(fakeSleeper, fakeSleeper)

		want := []string{
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}
		got := fakeSleeper.calls

		if diff := cmp.Diff(got, want); diff != "" {
			t.Fatalf("order mismatch: %s", diff)
		}
	})
}

type spyTime struct {
	durationSlept time.Duration
}

func (s *spyTime) Sleep(d time.Duration) {
	s.durationSlept = d
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &spyTime{}
	sleeper := configurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.sleep(sleepTime)

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
