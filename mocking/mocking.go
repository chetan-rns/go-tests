package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countDownStart = 3
	finalWord      = "Go!"
)

type Sleeper interface {
	Sleep()
}
type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	countDown(os.Stdout, sleeper)
}

func countDown(w io.Writer, sleeper Sleeper) {
	for i := countDownStart; i >= 1; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
	sleeper.Sleep()
	fmt.Fprint(w, finalWord)
}
