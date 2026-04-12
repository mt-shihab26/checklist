package main

import (
	"concurrency/order"
	"fmt"
)

func main() {
	orders := order.Generates(20)

	order.Process(orders)

	order.UpdateOrdersStatus(orders)

	order.ReportOrderStatus(orders)

	fmt.Println("All operations completed. Exiting.")
}
