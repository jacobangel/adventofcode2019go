package intcodecomp

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

func Get10() int {
	return 10
}

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
		cmd = cmd / 10
		value = value + divisions*digit
		divisions = divisions * 10
	}
	return value, cmd
}

func getInstructionWidth(code Opcode) int {
	return InstructionLen[code]
}

func parseInstruction(cmd int) (Opcode, []ParameterMode) {
	instruction, paramValue := takeDigits(cmd+1-1, 2)
	inst := Opcode(instruction)
	paramCommands, ok := getParameterModeList(inst, paramValue)
	if !ok {
		fmt.Printf("Error encountered! Opcode %d from %d (%d) gave %d\n", instruction, cmd, paramValue, inst)
		inst = ERROR
	}
	return inst, paramCommands
}

func getParameterModeList(inst Opcode, paramValue int) ([]ParameterMode, bool) {
	paramCount := getInstructionWidth(inst)
	if paramCount == 0 {
		return []ParameterMode{}, false
	}

	paramCommands := make([]ParameterMode, paramCount)
	for i := 0; i < paramCount; i++ {
		param, newParam := takeDigits(paramValue, 1)
		paramCommands[i] = ParameterMode(param)
		paramValue = newParam
	}
	return paramCommands, true
}

func getMemoryAddress(argument []int, mode []ParameterMode, data []int, index int) int {
	address := argument[index]
	if mode[index] == IMMEDIATE {
		return address
	}
	return data[address]
}

func makeArguments(parameterModes []ParameterMode, data []int) {

}

func InterpretProgram(data []int, input int) int {
	output := 0
	for instructionPointer := 0; instructionPointer < len(data); {
		controlChar := data[instructionPointer]
		instruction, readParamOpts := parseInstruction(controlChar)

		arguments := make([]int, len(readParamOpts))
		for i := 0; i < len(readParamOpts) && instructionPointer+i+1 < len(data); i++ {
			fmt.Printf("%d, IP: %d to %d\n", len(data), (instructionPointer), instruction)
			reference := instructionPointer + i + 1
			arguments[i] = data[reference]
		}

		// fmt.Printf("Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
		switch instruction {
		case ERROR:
			fmt.Printf("The program encountered an illegal error. Stack dump:\n%v \n", data[0:instructionPointer+1])
			return -1
		case ADD:
			a := getMemoryAddress(arguments, readParamOpts, data, 0)
			b := getMemoryAddress(arguments, readParamOpts, data, 1)
			// fmt.Printf("Adding %d + %d and storing in register %d\n", a, b, arguments[2])
			data[arguments[2]] = a + b
		case MULT:
			a := getMemoryAddress(arguments, readParamOpts, data, 0)
			b := getMemoryAddress(arguments, readParamOpts, data, 1)
			result := a * b
			data[arguments[2]] = result
		case GET:
			fmt.Printf("Control Get: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
			output = getMemoryAddress(arguments, readParamOpts, data, 0)
			fmt.Printf("\nPrinting output: %d\n", output)
		case STORE:
			fmt.Printf("Storing %d in register %d\n", input, arguments[0])
			data[arguments[0]] = input
		case JUMP_IF_FALSE:
			a := getMemoryAddress(arguments, readParamOpts, data, 0)
			b := getMemoryAddress(arguments, readParamOpts, data, 1)
			if a == 0 {
				fmt.Printf("JIF, Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
				instructionPointer = b
				continue
			}
		case JUMP_IF_TRUE:
			a := getMemoryAddress(arguments, readParamOpts, data, 0)
			b := getMemoryAddress(arguments, readParamOpts, data, 1)
			if a != 0 {
				fmt.Printf("JIT: Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
				instructionPointer = b
				continue
			}
		case LESS_THAN:
			a := getMemoryAddress(arguments, readParamOpts, data, 0)
			b := getMemoryAddress(arguments, readParamOpts, data, 1)
			result := a < b
			fmt.Printf("LT: Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
			if result {
				data[arguments[2]] = 1
			} else {
				data[arguments[2]] = 0
			}
		case EQUALS:
			a := getMemoryAddress(arguments, readParamOpts, data, 0)
			b := getMemoryAddress(arguments, readParamOpts, data, 1)
			result := a == b
			fmt.Printf("EQ: Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
			if result {
				data[arguments[2]] = 1
			} else {
				data[arguments[2]] = 0
			}
		case STOP:
			fmt.Println("The program has completed without error!")
			return output
		}
		instructionPointer = 1 + instructionPointer + len(arguments)
	}

	return 0
}
