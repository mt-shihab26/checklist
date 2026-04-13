package main

import (
	"fmt"
	"sync"
)

func main() {
	waitGroup := sync.WaitGroup{}
	mutex := sync.Mutex{}

	count := 0

	for range 1000 {
		waitGroup.Go(func() {
			mutex.Lock()
			count++
			fmt.Println("counting:", count)
			mutex.Unlock()
		})
	}

	waitGroup.Wait()

	fmt.Println("Counter:", count)
}
