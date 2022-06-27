package main

import (
	"fmt"
	"os"
	"strconv"
)

func pow(x, y int) (res int) {
	res = x
	for i := 1; i < y; i++ {
		res *= x
	}

	return
}

func task3() {
	if len(os.Args) != 3 {
		fmt.Println("Enter numbers")
		return
	}

	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil {
		fmt.Println("Provide valid args")
	}

	res := pow(x, y)
	fmt.Println(res)
}

func main() {
	task3()
}
