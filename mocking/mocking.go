package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type sleeper struct {
	duration time.Duration
}

func (s *sleeper) Sleep() {
	time.Sleep(s.duration)
}

const (
	startCount = 3
	finalWord  = "Go!"
)

func main() {
	countDown(os.Stdout, &sleeper{1 * time.Second})
}

func countDown(out io.Writer, sleeper Sleeper) {
	for i := startCount; i >= 1; i-- {
		sleeper.Sleep()
		out.Write([]byte(fmt.Sprintf("%d\n", i)))
	}
	sleeper.Sleep()

	out.Write([]byte(finalWord))
}
