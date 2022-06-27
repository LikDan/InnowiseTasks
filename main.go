package main

import (
	"fmt"
	"os"
	"strings"
)

func isVowel(input string) bool {
	return strings.IndexAny(input, "aeiouy") == 0
}

func task1() {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	if len(os.Args) != 2 {
		fmt.Println("Enter a letter")
		return
	}

	letter := os.Args[1]
	if len(letter) != 1 {
		fmt.Println("Not valid input")
		return
	}

	if strings.IndexAny(letter, alphabet) == -1 {
		fmt.Println("Not valid input")
		return
	}

	if isVowel(letter) {
		fmt.Printf("'%v' is a vowel\n", letter)
	} else {
		fmt.Printf("'%v' is a consonant\n", letter)
	}
}

func main() {
	task1()
}
