package main

import (
	"fmt"
)

func fibonacci(n, x, y int, ch chan int) {

	for ; n > 0; n-- {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	x, y := 0, 1
	ch := make(chan int, 10)
	go fibonacci(20, x, y, ch)

	for i := range ch {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
