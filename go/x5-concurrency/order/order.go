package order

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func Generates(count int) []*Order {
	orders := []*Order{}

	for i := range count {
		orders = append(orders, &Order{ID: i + 1, Status: "pending"})
	}

	return orders
}

func Process(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.IntN(500) * int(time.Millisecond)))

		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func UpdateOrdersStatus(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.IntN(500) * int(time.Millisecond)))

		status := []string{"processing", "shipped", "delivered"}[rand.IntN(3)]

		order.Status = status

		fmt.Printf("Updated order %d status: %s\n", order.ID, status)
	}
}

func ReportOrderStatus(orders []*Order) {
	for range 5 {
		time.Sleep(1 * time.Second)

		fmt.Printf("\n--- Order Status Report ---")

		for _, order := range orders {
			fmt.Printf("Order %d: %s\n", order.ID, order.Status)
		}

		fmt.Println("----------------------------")
	}
}
