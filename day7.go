package main

import (
	"fmt"
)

func day71() {
	fmt.Println("Day 7.1")

	programData := LoadProgramData("./day7_input.txt")
	fmt.Println(programData)
	InterpretProgram(programData, 1)
	fmt.Printf("Opcodes: %d, %d, %d, %d, %d\n", ADD, MULT, STORE, GET, STOP)
	fmt.Printf("The answer is: '%d'\n", 42)
}

func day72() {
	fmt.Println("Day 7.2")
	programData := LoadProgramData("./day7_input.txt")
	fmt.Println(programData)
	InterpretProgram(programData, 7)
	fmt.Printf("Opcodes: %d, %d, %d, %d, %d\n", ADD, MULT, STORE, GET, STOP)
	fmt.Printf("The answer is: '%d'\n", 42)
}

func day7() {
	fmt.Println("Day 7")
	fmt.Println("---------")
	day71()
	day72()
	fmt.Println()
}
