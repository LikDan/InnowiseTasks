package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func fibonacci(num int) int {
	if num < 0 {
		return 0
	}

	if num == 1 || num == 2 {
		return 1
	}

	return fibonacci(num-1) + fibonacci(num-2)
}

func fibonacciLoop(num int) int {
	if num < 1 {
		return 0
	}

	if num == 2 {
		return 1
	}

	n1 := 1
	n2 := 0

	for i := 1; ; i++ {
		if i == num {
			return n1
		}

		n := n1 + n2

		n2 = n1
		n1 = n
	}
}

func task5() {
	if len(os.Args) != 2 {
		fmt.Println("Enter a number")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	t := time.Now().UnixMicro()
	for i := 0; i < num; i++ {
		fmt.Print(fibonacci(i+1), "\t")
	}
	t = time.Now().UnixMicro() - t
	fmt.Println("Recursion: " + strconv.FormatInt(t, 10))

	t = time.Now().UnixMicro()
	for i := 0; i < num; i++ {
		fmt.Print(fibonacciLoop(i+1), "\t")
	}
	t = time.Now().UnixMicro() - t
	fmt.Println("Loop: " + strconv.FormatInt(t, 10))
}
