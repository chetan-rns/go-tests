package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countDownStart = 3
)

func main() {
	sleeper := &configurableSleeper{1 * time.Second, time.Sleep}
	countDown(os.Stdout, sleeper)
}

type sleeper interface {
	Sleep()
}

type defaultSleeper struct{}

func (d *defaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type configurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *configurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func countDown(out io.Writer, s sleeper) {
	for i := countDownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintf(out, "%d\n", i)
	}
	s.Sleep()
	fmt.Fprint(out, finalWord)
}
