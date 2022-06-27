package main

import (
	"fmt"
	"os"
	"strconv"
)

func pow(x, y int) (res int, err error) {
	if y == 0 {
		if x == 0 {
			return 0, fmt.Errorf("not valid data")
		}

		return 1, nil
	}

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

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Provide valid args")
	}

	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Provide valid args")
	}

	res, err := pow(x, y)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func main() {
	task3()
}
