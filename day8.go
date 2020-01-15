package main

import (
	"fmt"
)


func day81() {
	fmt.Println("Day 8.1")
	programData := LoadProgramData("./day8_input.txt")
	fmt.Println(programData)
}

func day82() {
	fmt.Println("Day 8.2")
	programData := LoadProgramData("./day8_input.txt")
	fmt.Println(programData)
}

func day8() {
	fmt.Println("Day 8")
	fmt.Println("---------")
	day81()
	day82()
	fmt.Println()
}
