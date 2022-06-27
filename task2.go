package main

import (
	"fmt"
	"os"
	"strings"
)

func isVowelSwitch(input string) bool {
	switch strings.IndexAny(input, "aeiouy") {
	case 0:
		return true
	default:
		return false
	}
}

func task2() {
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

	if isVowelSwitch(letter) {
		fmt.Printf("'%v' is a vowel\n", letter)
	} else {
		fmt.Printf("'%v' is a consonant\n", letter)
	}
}
