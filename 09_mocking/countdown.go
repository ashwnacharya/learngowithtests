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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (d ConfigurableSleeper) Sleep() {
	d.sleep(d.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{ duration: 1 * time.Second, sleep: time.Sleep }
	Countdown(os.Stdout, sleeper)
}