package main

import (
	"fmt"
)

const LOWER_BOUND = 235741
const UPPER_BOUND = 706948

func isValidGroupedPassword(guess int) bool {
	return hasEvenDouble(guess) && isAscending((guess))
}

func isValidPassword(guess int) bool {
	return hasDouble(guess) && isAscending(guess)
}

func hasEvenDouble(guess int) bool {

	if guess < 10 {
		return false
	}

	lastDigit := -1
	onaRun := false
	decremented := false
	doubleCount := 0
	for guess > 0 {
		nextDigit := guess % 10
		if nextDigit == lastDigit {
			if onaRun {
				if !decremented {
					doubleCount--
				}
				decremented = true
			} else {
				doubleCount += 1
				decremented = false
				onaRun = true
			}
		} else {
			onaRun = false
			decremented = false
		}

		lastDigit = nextDigit
		guess = guess / 10
	}
	return doubleCount != 0
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

func crackTheCode2() int {
	count := 0
	for i := LOWER_BOUND + 1; i < UPPER_BOUND; i++ {
		if isValidGroupedPassword(i) {
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

func day42() {
	fmt.Println("Day 4.2")
	code := crackTheCode2()
	fmt.Printf("The possible passwords are: '%d'\n", code)
}

func day4() {
	fmt.Println("Day 4")
	fmt.Println("---------")
	day41()
	day42()
	fmt.Println()
}
