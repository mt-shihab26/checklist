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
