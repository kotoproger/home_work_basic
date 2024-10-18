package main

import (
	"fmt"
	"sync"

	"github.com/kotoproger/home_work_basic/hw11_worker_pool/worker"
)

func main() {
	counter := 0

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker.Increment(&counter, &mu)
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
