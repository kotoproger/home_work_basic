package worker

import (
	"fmt"
	"sync"
)

func Increment(counter *int, mutex *sync.Mutex) {
	mutex.Lock()
	*counter++
	mutex.Unlock()

	fmt.Println("increment done")
}
