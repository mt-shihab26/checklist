package main

import (
	"fmt"
	"sync"
)

func main() {
	waitGroup := sync.WaitGroup{}
	mutex := sync.RWMutex{}

	data := 0

	waitGroup.Go(func() {
		for i := range 3 {
			mutex.Lock()
			data = i
			fmt.Println("Writing:", data)
			mutex.Unlock()
		}
	})

	for id := range 3 {
		waitGroup.Go(func() {
			for range 3 {
				mutex.RLock()
				fmt.Printf("Reader %d sees: %d\n", id, data)
				mutex.RUnlock()
			}

		})
	}

	waitGroup.Wait()

	fmt.Println("In Last Data:", data)
}
