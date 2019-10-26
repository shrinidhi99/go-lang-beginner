package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(1000)
}

func main() {
	fmt.Println("Hello from goroutine")
	go sayHello()
	fmt.Println("End of goroutine")
	sayHello()
}
