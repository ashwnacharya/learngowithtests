package countersync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()
		assertCounter(t, &counter, wantedCount)
	})

}

func assertCounter(t *testing.T, counter *Counter, want int) {
	if counter.Value() != want {
		t.Errorf("got %d want %d", counter.Value(), want)
	}
}