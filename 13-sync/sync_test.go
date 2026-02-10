package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Increment the count 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Int()
		counter.Int()
		counter.Int()

		assertCounter(t, counter, 3)
	})

	t.Run("It runs safely with concurrency", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				defer wg.Done()
				counter.Int()
			}()
		}

		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, counter *Counter, want int) {
	t.Helper()

	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}
