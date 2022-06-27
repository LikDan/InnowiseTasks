package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func FibonacciGoroutine(ch chan int, e chan bool) {
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

func task7() {
	if len(os.Args) != 2 {
		log.Fatal("Enter a number")
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal("Provide valid agrs")
	}

	ch := make(chan int)
	quit := make(chan bool)

	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-ch)
		}
		quit <- true
	}(n)

	FibonacciGoroutine(ch, quit)
}
