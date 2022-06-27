package main

import (
	"fmt"
	"os"
	"strconv"
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

	f := fibonacci(num)
	fmt.Println(f)
}

func main() {
	task5()
}
