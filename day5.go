package main

import (
	"fmt"
)

type Opcode int

const (
	ERROR Opcode = iota
	ADD
	MULT
	STORE
	GET
	JUMP_IF_TRUE
	JUMP_IF_FALSE
	LESS_THAN
	EQUALS
	STOP Opcode = 99
)

type ParameterMode int

const (
	POSITION ParameterMode = iota
	IMMEDIATE
)

var InstructionLen = map[Opcode]int{
	ADD:           3,
	MULT:          3,
	STORE:         1,
	GET:           1,
	STOP:          1,
	JUMP_IF_FALSE: 2,
	JUMP_IF_TRUE:  2,
	LESS_THAN:     3,
	EQUALS:        3,
}

func takeDigits(cmd int, count int) (int, int) {
	if cmd < 10 || count < 1 {
		return cmd, 0
	}
	divisions := 1
	value := 0
	for ; count > 0; count-- {
		digit := cmd % 10
		value = value + divisions*digit
		cmd = cmd / 10
		divisions = divisions * 10
	}
	return value, cmd
}

func getInstructionWidth(code int) int {
	return InstructionLen[Opcode(code)]
}

func parseInstruction(cmd int) (Opcode, []ParameterMode) {
	instruction, paramValue := takeDigits(cmd+1-1, 2)
	if instruction == 99 {
		return STOP, []ParameterMode{}
	}
	paramCount := getInstructionWidth(instruction)
	if paramCount == 0 {
		fmt.Printf("Illegal count encountered! Opcode %d from %d (%d) gave %d\n", instruction, cmd, paramValue, paramCount)
		return ERROR, []ParameterMode{}
	}
	paramCommands := make([]ParameterMode, paramCount)
	for i := 0; i < paramCount; i++ {
		param, newParam := takeDigits(paramValue, 1)
		paramCommands[i] = ParameterMode(param)
		paramValue = newParam
	}
	return Opcode(instruction), paramCommands
}

func handleReadMemory(reference []int, mode []ParameterMode, data []int, index int) int {
	address := reference[index]
	reversed := len(mode) - index - 1
	reversed = index
	if address < 0 || address > len(data)-1 {
		fmt.Printf("%d %v %v\n", address, reference, mode)
		return address
	}
	if mode[reversed] == IMMEDIATE {
		return address
	}
	return data[address]
}

func InterpretProgram(data []int, input int) int {
	for instructionPointer := 0; instructionPointer < len(data); {
		controlChar := data[instructionPointer]
		instruction, readParamOpts := parseInstruction(controlChar)

		arguments := make([]int, len(readParamOpts))
		for i := 0; i < len(readParamOpts); i++ {
			reference := instructionPointer + i + 1
			arguments[i] = data[reference]
		}

		// fmt.Printf("Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
		switch instruction {
		case ERROR:
			fmt.Printf("The program encountered an illegal error. Stack dump:\n%v \n", data[0:instructionPointer+1])
			return 0
		case ADD:
			a := handleReadMemory(arguments, readParamOpts, data, 0)
			b := handleReadMemory(arguments, readParamOpts, data, 1)
			fmt.Printf("Adding %d + %d and storing in register %d\n", a, b, arguments[2])
			data[arguments[2]] = a + b
		case MULT:
			a := handleReadMemory(arguments, readParamOpts, data, 0)
			b := handleReadMemory(arguments, readParamOpts, data, 1)
			result := a * b
			data[arguments[2]] = result
		case GET:
			fmt.Printf("Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
			fmt.Printf("\nPrinting output: %d\n", data[arguments[0]])
		case STORE:
			fmt.Printf("Storing %d in register %d\n", input, arguments[0])
			data[arguments[0]] = input
		case JUMP_IF_FALSE:
			a := handleReadMemory(arguments, readParamOpts, data, 0)
			b := handleReadMemory(arguments, readParamOpts, data, 1)
			if a == 0 {
				fmt.Printf("JIF, Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
				instructionPointer = b
				continue
			}
		case JUMP_IF_TRUE:
			a := handleReadMemory(arguments, readParamOpts, data, 0)
			b := handleReadMemory(arguments, readParamOpts, data, 1)
			if a != 0 {
				fmt.Printf("JIT: Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
				instructionPointer = b
				continue
			}
		case LESS_THAN:
			a := handleReadMemory(arguments, readParamOpts, data, 0)
			b := handleReadMemory(arguments, readParamOpts, data, 1)
			result := a < b
			fmt.Printf("LT: Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
			if result {
				data[arguments[2]] = 1
			} else {
				data[arguments[2]] = 0
			}
		case EQUALS:
			a := handleReadMemory(arguments, readParamOpts, data, 0)
			b := handleReadMemory(arguments, readParamOpts, data, 1)
			result := a == b
			fmt.Printf("EQ: Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
			if result {
				data[arguments[2]] = 1
			} else {
				data[arguments[2]] = 0
			}
		case STOP:
			fmt.Println("The program has completed without error!")
			return 0
		}
		instructionPointer = 1 + instructionPointer + len(arguments)
	}

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
	InterpretProgram(programData, 1)
	fmt.Printf("Opcodes: %d, %d, %d, %d, %d\n", ADD, MULT, STORE, GET, STOP)
	fmt.Printf("The answer is: '%d'\n", 42)
}

func day52() {
	fmt.Println("Day 5.2")
	programData := LoadProgramData("./day5_input.txt")
	fmt.Println(programData)
	InterpretProgram(programData, 5)
	fmt.Printf("Opcodes: %d, %d, %d, %d, %d\n", ADD, MULT, STORE, GET, STOP)
	fmt.Printf("The answer is: '%d'\n", 42)
}

func day5() {
	fmt.Println("Day 5")
	fmt.Println("---------")
	day51()
	day52()
	fmt.Println()
}
