package main

import (
	"fmt"
)

const LOWER_BOUND = 235741
const UPPER_BOUND = 706948

func isValidPassword(guess int) bool {
	return hasDouble(guess) && isAscending(guess)
}

func hasDouble(guess int) bool {
	if guess < 10 {
		return false
	}
	lastDigit := -1
	for guess > 0 {
		nextDigit := guess % 10
		if nextDigit == lastDigit {
			return true
		}
		lastDigit = nextDigit
		guess = guess / 10
	}
	return false
}

func isAscending(guess int) bool {
	if guess < 10 {
		return true
	}
	lastDigit := -1
	for guess > 0 {
		nextDigit := guess % 10
		if nextDigit > lastDigit && lastDigit != -1 {
			return false
		}
		lastDigit = nextDigit
		guess = guess / 10
	}

	return true
}

func crackTheCode() int {
	count := 0
	for i := LOWER_BOUND; i <= UPPER_BOUND; i++ {
		if isValidPassword(i) {
			count++
		}
	}
	return count
}

func day41() {
	fmt.Println("Day 4.1")
	code := crackTheCode()
	fmt.Printf("The possible passwords are: '%d'\n", code)
}

func day4() {
	fmt.Println("Day 4")
	fmt.Println("---------")
	day41()
	fmt.Println()
}
