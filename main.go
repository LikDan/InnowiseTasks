package main

import (
	"fmt"
)

func fibonacci(ch chan int, e chan bool) {
	n1 := 1
	n2 := 0

	ch <- 1

	for {
		n := n1 + n2

		n2 = n1
		n1 = n

		select {
		case ch <- n:
		case _ = <-e:
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	n := 10

	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-ch)
		}
		quit <- true
	}(n)

	fibonacci(ch, quit)
}
