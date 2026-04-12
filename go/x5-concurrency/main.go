package main

import (
	"concurrency/order"
	"fmt"
	"sync"
)

func main() {
	orders := order.Generates(100)

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(3)

	go func() {
		defer waitGroup.Done()
		order.Process(orders)
	}()

	go func() {
		defer waitGroup.Done()
		order.UpdateOrdersStatus(orders)
	}()

	go func() {
		defer waitGroup.Done()
		order.ReportOrderStatus(orders)
	}()

	waitGroup.Wait()

	fmt.Println("All operations completed. Exiting.")
}
