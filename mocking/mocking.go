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
	countDown(os.Stdout, &defaultSleeper{})
}

type sleeper interface {
	Sleep()
}

type defaultSleeper struct{}

func (d *defaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func countDown(out io.Writer, s sleeper) {
	for i := countDownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintf(out, "%d\n", i)
	}
	s.Sleep()
	fmt.Fprint(out, finalWord)
}
