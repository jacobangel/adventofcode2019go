package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

func ConvertToSlice(str string) []int {
	if len(str) == 0 {
		return []int{}
	}
	retValue := []int{}
	for index, item := range strings.Split(str, ",") {
		intVal, err := strconv.Atoi(item)
		if err != nil {
			fmt.Printf("Could not convert value '%s' of bucket %d to str. %g", item, index, err)
			return retValue
		}
		// fmt.Printf("-- value: %d, index: %d\n", intVal, index)
		retValue = append(retValue, intVal)
	}
	return retValue
}

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
