package intcodecomp

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
		cmd = cmd / 10
		value = value + divisions*digit
		divisions = divisions * 10
	}
	return value, cmd
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
	paramCount := InstructionLen[inst]
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

func resolveValueFromMemory(address int, mode ParameterMode, data []int) int {
	if mode == IMMEDIATE {
		return address
	}
	return data[address]
}

func gatherArguments(parameterModes []ParameterMode, instructionPointer int, data []int) []int {
	arguments := make([]int, len(parameterModes))
	for i := 0; i < len(parameterModes) && instructionPointer+i+1 < len(data); i++ {
		reference := instructionPointer + i + 1
		arguments[i] = data[reference]
	}

	return arguments
}

func eq(args []int) int {
	if args[0] == args[1] {
		return 1
	}
	return 0
}

func lt(args []int) int {
	if args[0] < args[1] {
		return 1
	}
	return 0
}

func add(args []int) int {
	return args[1] + args[0]
}

func mult(args []int) int {
	return args[1] * args[0]
}

func operate(operation func([]int) int, rawArgs []int, readOpts []ParameterMode, data []int) {
	arguments := make([]int, len(rawArgs))
	for index, item := range rawArgs {
		arguments[index] = resolveValueFromMemory(item, readOpts[index], data)
	}
	result := operation(arguments)
	// this is ugly!
	store(data, rawArgs, result)
}

func store(data []int, args []int, value int) {
	destAddr := args[len(args)-1]
	data[destAddr] = value
}

/**
Each instruction has a series of phases.
1. operation, get value
2. store
3. jump (usually width)

*/
func InterpretProgram(data []int, input []int, startingPointer int) (int, Opcode, []int, int) {
	output := 0
	readInput := func() func() (int, bool) {
		counter := -1
		return func() (int, bool) {
			counter += 1
			if counter >= len(input) {
				return -1, false
			}
			return input[counter], true
		}
	}()

	for instructionPointer := startingPointer; instructionPointer < len(data); {
		controlChar := data[instructionPointer]
		instruction, readParamOpts := parseInstruction(controlChar)
		arguments := gatherArguments(readParamOpts, instructionPointer, data)

		// handle operations with work
		switch instruction {
		case ERROR:
			fmt.Printf("The program encountered an illegal error. Stack dump:\n%v \n", data[0:instructionPointer+1])
			return -1, ERROR, data, instructionPointer

		case ADD:
			operate(add, arguments, readParamOpts, data)

		case MULT:
			operate(mult, arguments, readParamOpts, data)

		case GET:
			output = resolveValueFromMemory(arguments[0], readParamOpts[0], data)
			// do i need  to return the increment?
			return output, GET, data, instructionPointer + len(arguments) + 1
			// fmt.Printf("\nPrinting output: %d\n", output)

		case STORE:
			toRead, ok := readInput()
			if !ok {
				return output, GET, data, instructionPointer // + len(arguments) + 1
			}
			store(data, arguments, toRead)

		case LESS_THAN:
			operate(lt, arguments, readParamOpts, data)

		case EQUALS:
			operate(eq, arguments, readParamOpts, data)

		case STOP:
			// fmt.Println("The program has completed without error!")
			return data[len(data) - 2], STOP, data, instructionPointer
		}

		// handle jumps
		switch instruction {
		case JUMP_IF_FALSE:
			a := resolveValueFromMemory(arguments[0], readParamOpts[0], data)
			b := resolveValueFromMemory(arguments[1], readParamOpts[1], data)
			if a == 0 {
				// fmt.Printf("JIF, Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
				instructionPointer = b
				continue
			}

		case JUMP_IF_TRUE:
			a := resolveValueFromMemory(arguments[0], readParamOpts[0], data)
			b := resolveValueFromMemory(arguments[1], readParamOpts[1], data)
			if a != 0 {
				// fmt.Printf("JIT: Control: %d, Inst: %d, readOps: %v, args: %d, index: %d\n", controlChar, instruction, readParamOpts, arguments, instructionPointer)
				instructionPointer = b
				continue
			}
		}
		instructionPointer += 1 + len(arguments)
	}

	return -1, ERROR, data, 0
}
