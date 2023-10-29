package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpyWriter struct {
	Buffer bytes.Buffer
	CallLog *[]string
}

func (s *SpyWriter) Write(p []byte) (n int, err error) {
	*s.CallLog = append(*s.CallLog, write)
	return s.Buffer.Write(p)
}

type SpySleeper struct {
	Calls int
	CallLog *[]string
}

func (s *SpySleeper) Sleep() {
	s.Calls++
	*s.CallLog = append(*s.CallLog, sleep)
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {

	got_callLog := []string{}

	spyWriter := SpyWriter{Buffer: bytes.Buffer{}, CallLog: &got_callLog}
	spySleeper := SpySleeper{CallLog: &got_callLog}

	Countdown(&spyWriter, &spySleeper)

	want_callLog := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !reflect.DeepEqual(want_callLog, got_callLog) {
		t.Errorf("got %q want %q", got_callLog, want_callLog)
	}


	got := spyWriter.Buffer.String()
	want:= "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls, want 3, got %d", spySleeper.Calls)
	}
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
