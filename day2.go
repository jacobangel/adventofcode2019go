/**
Opcodes:
1 - read the next two positions addresses, add them and then store them in the third position.
2 - read the next two positions, multiply them, and then store in the third
99 - Termination of the program.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const stub = 10.0

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

func RunProgram(data []int) int {
	// read slice of 4
	// if we hit 99 in the first slot of the slice, then exit.
	//
	for i := 0; i < len(data); i += 4 {
		controlChar := data[i]
		if controlChar == 99 {
			return data[0]
		}

		a := data[data[i+1]]
		b := data[data[i+2]]
		register := data[i+3]

		if controlChar == 1 {
			data[register] = a + b
		}
		if controlChar == 2 {
			data[register] = a * b
		}
		//fmt.Printf("%d %d %d %d\n", controlChar, a, b, register)
		// fmt.Printf("%d\n", data)
	}

	return 0
}

func day1Part1() {
	// load up the file
	// read them into a slice.
	fileInput := loadFile("./day2-input.txt")
	programData := ConvertToSlice(fileInput)

	// replace the value in position 1 with 12
	programData[1] = 12
	// and replace the value in position 2 with 2
	programData[2] = 2
	// and run the program.
	fmt.Printf("The program data is now: %d", programData)
	firstResult := RunProgram(programData)

	fmt.Printf("The value of the result is: %d\n", firstResult)
	// print out the result of position 0!
	// end!
}

func day2Part2() {
	fmt.Println("Starting Part 2!")
	fileInput := loadFile("./day2-input.txt")
	valueToFind := 19690720

	for nounIndex := 0; nounIndex < 100; nounIndex++ {
		for verbIndex := 0; verbIndex < 100; verbIndex++ {
			programData := ConvertToSlice(fileInput)
			programData[1] = nounIndex
			programData[2] = verbIndex
			value := RunProgram(programData)
			if value == valueToFind {
				fmt.Printf("Found it!: %d\n", nounIndex*100+verbIndex)
			}
		}
	}
	fmt.Printf("No result found!\n")
}

func day2() {
	day1Part1()
	day2Part2()
}
