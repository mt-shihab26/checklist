package main

import (
	"concurrency/order"
	"fmt"
)

func main() {
	orders := order.Generates(20)

	order.Process(orders)

	fmt.Println("All operations completed. Exiting.")
}
