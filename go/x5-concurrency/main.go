package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	orders := generatesOrders(20)
	waitGroup := sync.WaitGroup{}
	waitGroup.Go(func() {
		updateOrdersStatus(1, orders)
	})
	waitGroup.Go(func() {
		updateOrdersStatus(2, orders)
	})
	waitGroup.Go(func() {
		updateOrdersStatus(3, orders)
	})
	waitGroup.Wait()
	reportOrderStatus(orders)
	fmt.Println("All operations completed. Exiting.")
}

type Order struct {
	ID     int
	Status string
	Mutex  sync.Mutex
}

func generatesOrders(count int) []*Order {
	orders := []*Order{}
	for i := range count {
		orders = append(orders, &Order{ID: i + 1, Status: "pending"})
	}
	return orders
}

func updateOrdersStatus(no int, orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.IntN(500) * int(time.Millisecond)))
		status := []string{"processing", "shipped", "delivered"}[rand.IntN(3)]
		order.Mutex.Lock()
		order.Status = status
		order.Mutex.Unlock()
		fmt.Printf("Routine #%d: Updated order %d status: %s\n", no, order.ID, status)
	}
}

func reportOrderStatus(orders []*Order) {
	time.Sleep(1 * time.Second)
	fmt.Printf("--- Order Status Report ---\n")
	for _, order := range orders {
		fmt.Printf("Order %d: %s\n", order.ID, order.Status)
	}
	fmt.Println("----------------------------")
}
