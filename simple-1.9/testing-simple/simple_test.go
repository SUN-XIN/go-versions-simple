package p

import (
	"testing"
	"time"
)

func failure(t *testing.T) {
	t.Helper()         // go1.9 -> This call silences this function in error reports.
	t.Fatal("failure") // go1.8 -> the error message printed here
}

func TestFailure(t *testing.T) {
	failure(t) // go1.9 -> the error message printed here
}

func TestFailureGoroutine(t *testing.T) {
	go func() {
		failure(t) // go1.9 -> the error message printed here
	}()

	time.Sleep(time.Second)
}
