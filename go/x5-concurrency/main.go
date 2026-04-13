package main

import (
	"concurrency/order"
	"fmt"
	"sync"
)

func main() {
	orders := order.Generates(20)
	waitGroup := sync.WaitGroup{}
	waitGroup.Go(func() {
		order.UpdateOrdersStatus(1, orders)
	})
	waitGroup.Go(func() {
		order.UpdateOrdersStatus(2, orders)
	})
	waitGroup.Go(func() {
		order.UpdateOrdersStatus(3, orders)
	})
	waitGroup.Wait()
	order.ReportOrderStatus(orders)
	fmt.Println("All operations completed. Exiting.")
}
