package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func loadFile(str string) string {
	if len(str) == 0 {
		return ""
	}

	dat, err := ioutil.ReadFile(str)
	hasErr := err != nil
	if hasErr {
		fmt.Printf("There was an error: %g\n", err)
		return ""
	}

	return strings.TrimSpace(string(dat))
}


func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}