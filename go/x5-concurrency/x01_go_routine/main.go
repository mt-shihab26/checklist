package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := range 3 {
		fmt.Println("hello", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go sayHello()

	for i := range 3 {
		fmt.Println("Main", i)
		time.Sleep(1 * time.Second)
	}
}
