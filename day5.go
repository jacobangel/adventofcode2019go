package main

import "fmt"

func InterpretProgram(data []int) int {
	return 0
}

func LoadProgramData(fileName string) []int {
	fileInput := loadFile(fileName)
	programData := ConvertToSlice(fileInput)
	return programData
}

func day51() {
	fmt.Println("Day 5.1")

	programData := LoadProgramData("./day5_input.txt")
	fmt.Println(programData)
	fmt.Printf("The answer is: '%d'\n", 42)
}

func day52() {
	fmt.Println("Day 5.2")
	fmt.Printf("The answer is: '%d'\n", 42)
}

func day5() {
	fmt.Println("Day 5")
	fmt.Println("---------")
	day51()
	day52()
	fmt.Println()
}
