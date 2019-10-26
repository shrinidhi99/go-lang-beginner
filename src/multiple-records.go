package main

import "fmt"

func swap(x int, y int) (int, int) {
	return y, x
}
func main() {
	fmt.Println(3, 4)
	fmt.Println("After swapping")
	fmt.Println(swap(3, 4))
}
