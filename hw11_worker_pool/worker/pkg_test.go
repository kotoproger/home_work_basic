package worker

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIncrementLocked(t *testing.T) {
	counter := 0
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(1)

	mutex.Lock()
	go func() {
		defer wg.Done()
		Increment(&counter, &mutex)
	}()

	time.Sleep(time.Microsecond * 10)
	assert.Equal(t, 0, counter)

	mutex.Unlock()
	wg.Wait()
	assert.Equal(t, 1, counter)

	assert.True(t, mutex.TryLock())
}

func TestIncrementUnLocked(t *testing.T) {
	counter := 0
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		Increment(&counter, &mutex)
	}()

	wg.Wait()

	assert.Equal(t, 1, counter)

	assert.True(t, mutex.TryLock())
}
