package main

import (
	"bytes"
	"reflect"
	"testing"
)

type mockSleeper struct {
	Calls int
}

func (m *mockSleeper) Sleep() {
	m.Calls++
}

type CountDownOperationsSpy struct {
	Calls []string
}

func (c *CountDownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountDownOperationsSpy) Write(b []byte) (int, error) {
	c.Calls = append(c.Calls, write)
	return 0, nil
}

const (
	sleep = "sleep"
	write = "write"
)

func TestCountDown(t *testing.T) {
	buff := &bytes.Buffer{}
	fake := &CountDownOperationsSpy{}
	countDown(buff, fake)
	want := `3
2
1
Go!`

	if buff.String() != want {
		t.Errorf("got %s, want %s", buff.String(), want)
	}

	t.Run("Checking the order of operations", func(t *testing.T) {
		fake := &CountDownOperationsSpy{}
		countDown(&bytes.Buffer{}, fake)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if reflect.DeepEqual(fake.Calls, want) {
			t.Errorf("got %v, want %v", fake.Calls, want)
		}
	})
}
