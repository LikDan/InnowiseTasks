package main

import (
	"fmt"
	"math"
	"os"
)

func toInt(str string) (i int, err error) {
	for index, b := range str {
		if int(b)-48 < 0 || int(b)-48 > 9 {
			return 0, fmt.Errorf("not a number")
		}

		if len(str) == index+1 {
			i += int(b) - 48
			break
		}
		i += (int(b) - 48) * int(math.Pow(10, float64(len(str)-index-1)))
	}

	return
}

func task4() {
	if len(os.Args) != 2 {
		fmt.Println("Enter a number")
		return
	}

	i, err := toInt(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(i)
}

func main() {
	task4()
}
