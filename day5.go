package main

import "fmt"

type Opcode int
const (
	ERROR Opcode = iota
	ADD 
	MULT
	STORE
	GET
	STOP Opcode = 99
) 

type ParameterMode int
const (
	 POSITION ParameterMode = iota
	 IMMEDIATE
)

var InstructionLen = map[Opcode]int {
	ADD: 4,
	MULT: 4,
	STORE: 1,
	GET: 1,
	STOP: 1,
} 

func takeDigits(cmd int, count int) (int, int) {
	if cmd < 10 {
		return cmd, 0
	}
	divisions := 0;
	value := 0;
	for cmd > 0 {
		digit := cmd % 10
		cmd = cmd / 10
		divisions += 1
		value = value + divisions * digit
	}
	return value, cmd
}

func getInstructionWidth(code int) int {
  return InstructionLen[Opcode(code)]
}

func parseInstruction(cmd int) (Opcode, []ParameterMode) {
	instruction, paramValue := takeDigits(cmd, 2)
	paramCount := getInstructionWidth(instruction);
	if paramCount == 0 {
		fmt.Printf("Illegal count encountered! Opcode %d gave %d\n", instruction, paramCount)
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

func handleReadMemory(reference int, mode ParameterMode, data []int) int {
	if mode == IMMEDIATE {
		return reference
	}
	return data[reference]
}

func InterpretProgram(data []int) int {
	for instructionPointer := 0; instructionPointer < len(data); {
		controlChar := data[instructionPointer];
		instruction, readParamOpts := parseInstruction(controlChar)

		arguments := make([]int, len(readParamOpts))
		for i := 0; i < len(readParamOpts); i++ {
			reference := instructionPointer + i + 1
			arguments[i] = handleReadMemory(reference, readParamOpts[i], data)
		}	

		fmt.Printf("Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
		switch instruction {
		case ERROR:
			fmt.Printf("The program encountered an illegal error. Stack dump:\n%v \n", data[0:instructionPointer])
				return 0
		case ADD:
			result := arguments[0] + arguments[1]
			data[arguments[2]] = result;
		case MULT:
	  	result := arguments[0] * arguments[1]
			data[arguments[2]] = result;
		case GET:
			data[arguments[0]] = 1337
		case STORE:
			data[arguments[0]] = 7331
		case STOP:
			fmt.Println("The program has completed without error!")
			return 0
		}
		instructionPointer = 1 + instructionPointer + getInstructionWidth(controlChar)
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
  InterpretProgram(programData)
  fmt.Printf("Opcodes: %d, %d, %d, %d, %d\n", ADD, MULT, STORE, GET, STOP) 
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
