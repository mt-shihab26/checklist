package main

import (
	"concurrency/order"
	"fmt"
	"sync"
)

func main() {
	orders := order.Generates(100)

	waitGroup := sync.WaitGroup{}

	waitGroup.Go(func() {
		order.Process(orders)
	})

	waitGroup.Go(func() {
		order.UpdateOrdersStatus(orders)
	})

	waitGroup.Wait()

	order.ReportOrderStatus(orders)

	fmt.Println("All operations completed. Exiting.")
}
